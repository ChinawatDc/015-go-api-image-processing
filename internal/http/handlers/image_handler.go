package handlers

import (
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"

	"github.com/ChinawatDc/015-go-api-image-processing/internal/imageproc"
)

type ImageHandler struct {
	Processor     *imageproc.Processor
	OriginalDir   string
	ProcessedDir  string
	StaticURL     string
	WatermarkText string
	WatermarkSize float64
}

func (h *ImageHandler) UploadAndProcess(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "file required"})
		return
	}

	srcPath := filepath.Join(h.OriginalDir, file.Filename)
	dstPath := filepath.Join(h.ProcessedDir, file.Filename)

	_ = c.SaveUploadedFile(file, srcPath)

	img, err := imaging.Open(srcPath)
	if err != nil {
		c.JSON(500, gin.H{"error": "invalid image"})
		return
	}

	resized := h.Processor.Resize(img, 800)
	watermarked, err := h.Processor.WatermarkText(resized, h.WatermarkText, h.WatermarkSize)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	_ = imaging.Save(watermarked, dstPath)

	c.JSON(200, gin.H{
		"success": true,
		"url":     h.StaticURL + "/processed/" + file.Filename,
	})
}
