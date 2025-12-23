# 015-go-api-image-processing

แลปนี้เป็นส่วนหนึ่งของซีรีส์ **Go API Course**  
หัวข้อ: **จัดการรูปภาพ (Image Processing) ด้วย Go – Resize & Watermark**

---

## 🎯 เป้าหมายของแลป

- อัปโหลดรูปภาพ (JPEG / PNG)
- Resize รูปภาพโดยรักษาสัดส่วน
- ใส่ Watermark (ข้อความ) บนรูปภาพ
- บันทึกรูปที่ประมวลผลแล้ว
- เปิดดูรูปผ่าน Static URL

---

## 🧱 Tech Stack

- Go
- Gin Framework
- Image Processing (`github.com/disintegration/imaging`)
- Font & Text Watermark (`golang.org/x/image/font`)
- Local File Storage

---

## 📁 โครงสร้างโปรเจกต์

```
.
├─ cmd/api/main.go
├─ internal/
│  ├─ config/
│  ├─ http/handlers/
│  ├─ imageproc/
│  └─ utils/
├─ assets/
│  └─ fonts/
├─ uploads/
│  ├─ original/
│  └─ processed/
├─ .env
└─ README.md
```

---

## ⚙️ Environment Variables (.env)

```env
APP_PORT=8080

UPLOAD_ORIGINAL_DIR=uploads/original
UPLOAD_PROCESSED_DIR=uploads/processed

MAX_UPLOAD_MB=10
ALLOWED_EXT=jpg,jpeg,png

STATIC_URL=/static

WATERMARK_TEXT=© Go API Course
WATERMARK_FONT_SIZE=24
```

---

## ▶️ วิธีรันโปรเจกต์

```bash
go run cmd/api/main.go
```

Server:
http://localhost:8080

---

## 🔐 API Endpoint

### Upload + Image Processing

POST /image/process  
Form field: file

---

## 🌐 Static Access

http://localhost:8080/static/processed/<filename>

---

## 🧠 Key Concepts

- Image processing pipeline
- Aspect ratio resize
- Text watermark with TTF font
- Clean separation of concerns
- Ready for S3 / MinIO

---

## 🚀 Next Steps

- Logo watermark
- Dynamic resize via query
- Thumbnail generator
- EXIF stripping
- JWT protected upload

---

MIT License
# 015-go-api-image-processing
