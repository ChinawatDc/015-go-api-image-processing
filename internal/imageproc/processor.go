package imageproc

import (
	"image"
	"image/color"
	"os"

	"github.com/disintegration/imaging"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

type Processor struct {
	FontPath string
}

func New(fontPath string) *Processor {
	return &Processor{FontPath: fontPath}
}

func (p *Processor) Resize(src image.Image, width int) image.Image {
	return imaging.Resize(src, width, 0, imaging.Lanczos)
}

func (p *Processor) WatermarkText(img image.Image, text string, size float64) (image.Image, error) {
	rgba := imaging.Clone(img)

	fontBytes, err := os.ReadFile(p.FontPath)
	if err != nil {
		return nil, err
	}

	ft, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(ft, &opentype.FaceOptions{
		Size: size,
		DPI:  72,
	})
	if err != nil {
		return nil, err
	}

	d := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 180}),
		Face: face,
	}

	b := rgba.Bounds()
	d.Dot = fixed.P(b.Max.X-300, b.Max.Y-20)
	d.DrawString(text)

	return rgba, nil
}
