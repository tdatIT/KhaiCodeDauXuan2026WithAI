# Random Fortune App

Ung dung Go su dung Fiber framework, tra ve loi chuc Tet ngau nhien bang Tieng Viet.

## API Endpoints

| Method | Path       | Mo ta                          |
|--------|------------|--------------------------------|
| GET    | `/`        | Tra ve loi chao nam moi        |
| GET    | `/fortune` | Tra ve loi chuc Tet ngau nhien |

## Yeu cau

- Go 1.25+
- Docker & Docker Compose

## Chay local

```bash
go mod download
go run main.go
```

Ung dung chay tai `http://localhost:3000`.

## Chay voi Docker

### Build va chay ung dung

```bash
docker build -t random-fortune-app .
docker run -p 3000:3000 random-fortune-app
```

### Chay voi Docker Compose (bao gom Redis)

```bash
docker compose up -d
```

## Cau truc du an

```
random-fortune-app/
├── main.go              # Entry point
├── go.mod               # Go module
├── go.sum               # Go dependencies checksum
├── Dockerfile           # Multi-stage Docker build
├── docker-compose.yml   # Docker Compose (Redis)
└── README.md
```
