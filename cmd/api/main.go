package main

import (
	"os"

	"github.com/ChinawatDc/015-go-api-image-processing/internal/config"
	"github.com/ChinawatDc/015-go-api-image-processing/internal/http/handlers"
	"github.com/ChinawatDc/015-go-api-image-processing/internal/imageproc"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	_ = os.MkdirAll(cfg.UploadOriginal, 0755)
	_ = os.MkdirAll(cfg.UploadProcessed, 0755)

	r := gin.Default()
	r.MaxMultipartMemory = cfg.MaxUploadMB << 20

	r.Static(cfg.StaticURL+"/original", cfg.UploadOriginal)
	r.Static(cfg.StaticURL+"/processed", cfg.UploadProcessed)

	proc := imageproc.New("assets/fonts/Kanit-Regular.ttf")

	h := &handlers.ImageHandler{
		Processor:     proc,
		OriginalDir:   cfg.UploadOriginal,
		ProcessedDir:  cfg.UploadProcessed,
		StaticURL:     cfg.StaticURL,
		WatermarkText: cfg.WatermarkText,
		WatermarkSize: cfg.WatermarkFontSize,
	}

	r.POST("/image/process", h.UploadAndProcess)

	r.Run(":" + cfg.AppPort)
}
