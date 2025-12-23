package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	UploadOriginal string
	UploadProcessed string

	MaxUploadMB int64
	AllowedExt  map[string]bool

	StaticURL string

	WatermarkText string
	WatermarkFontSize float64
}

func Load() Config {
	_ = godotenv.Load(".env")

	return Config{
		AppPort: os.Getenv("APP_PORT"),

		UploadOriginal: os.Getenv("UPLOAD_ORIGINAL_DIR"),
		UploadProcessed: os.Getenv("UPLOAD_PROCESSED_DIR"),

		MaxUploadMB: mustInt64("MAX_UPLOAD_MB", 10),
		AllowedExt: parseExt(os.Getenv("ALLOWED_EXT")),

		StaticURL: os.Getenv("STATIC_URL"),

		WatermarkText: os.Getenv("WATERMARK_TEXT"),
		WatermarkFontSize: mustFloat("WATERMARK_FONT_SIZE", 24),
	}
}

func parseExt(s string) map[string]bool {
	m := map[string]bool{}
	for _, e := range strings.Split(s, ",") {
		m[strings.TrimSpace(e)] = true
	}
	return m
}

func mustInt64(k string, d int64) int64 {
	v, _ := strconv.ParseInt(os.Getenv(k), 10, 64)
	if v == 0 {
		return d
	}
	return v
}

func mustFloat(k string, d float64) float64 {
	v, _ := strconv.ParseFloat(os.Getenv(k), 64)
	if v == 0 {
		return d
	}
	return v
}
