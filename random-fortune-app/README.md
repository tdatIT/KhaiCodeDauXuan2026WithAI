# 🎊 Random Fortune App - Khai Xuân Đầu Năm 2026

## 🌟 Giới thiệu

Dự án **Random Fortune App** là một phần của **Code Khai Xuân Đầu Năm 2026** - một dự án biểu tượng mang đến **hy vọng khỏi đầu năm 2026** với tâm niệm công việc được suông sẻ trong kỷ nguyên phát triển cũng AI.

Ứng dụng Go sử dụng **Fiber framework**, trả về lời chúc Tết ngẫu nhiên bằng Tiếng Việt, kết hợp công nghệ hiện đại để tạo ra những trải nghiệm tươi mới trong năm mới.

### 💡 Tôn chỉ

- **Khỏi đầu**: Bắt đầu năm với tâm thế tích cực, đầy năng lượng
- **Suông sẻ**: Công việc trôi chảy, hiệu quả với sự hỗ trợ của AI
- **Phát triển**: Tiếp tục học tập, cải thiện, đón đầu công nghệ

## 📡 API Endpoints

| Method | Path       | Mo ta                          |
|--------|------------|--------------------------------|
| GET    | `/`        | Tra ve loi chao nam moi 2026   |
| GET    | `/fortune` | Tra ve loi chuc Tet ngau nhien |

## ⚙️ Yêu cầu

- Go 1.25+
- Docker & Docker Compose

## 🚀 Chạy local

```bash
go mod download
go run main.go
```

Ứng dụng chạy tại `http://localhost:3000`.

## 🐳 Chạy với Docker

### Build và chạy ứng dụng

```bash
docker build -t random-fortune-app .
docker run -p 3000:3000 random-fortune-app
```

### Chạy với Docker Compose (bao gồm Redis)

```bash
docker compose up -d
```

## 📁 Cấu trúc dự án

```
random-fortune-app/
├── main.go              # Entry point
├── go.mod               # Go module
├── go.sum               # Go dependencies checksum
├── Dockerfile           # Multi-stage Docker build
├── docker-compose.yml   # Docker Compose (Redis)
└── README.md
```

## 🎉 Lời chúc Năm 2026

Dự án này được tạo ra với tâm niệm:

> *"Năm 2026 - Khỏi đầu với hy vọng, suông sẻ với AI, phát triển cùng công nghệ"*

Mong rằng mỗi lần sử dụng ứng dụng này, bạn sẽ nhận được những lời chúc ý nghĩa và cảm nhận được năng lượng tích cực cho năm mới.

---

**Được phát triển với 💜 và AI trong Năm 2026**
