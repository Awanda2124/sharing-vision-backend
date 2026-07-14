# Sharing Vision Backend Technical Test

Backend REST API untuk mengelola artikel (Post Article) menggunakan Golang, Gin, GORM, dan MySQL.

## Tech Stack

- Go 1.24+
- Gin (web framework)
- GORM (ORM)
- MySQL
- golang-migrate

## How to Run (Local)

### 1. Clone repository

```bash
git clone https://github.com/Awanda2124/sharing-vision-backend.git
cd sharing-vision-backend
```

### 2. Install dependency

```bash
go mod tidy
```

### 3. Create database

```sql
CREATE DATABASE article;
```

### 4. Copy environment

Copy `.env.example` menjadi `.env`, lalu sesuaikan konfigurasi database:
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=article
APP_PORT=8080

### 5. Run migration

```bash
migrate -path migrations -database "mysql://root:@tcp(127.0.0.1:3306)/article" up
```

### 6. Run application

```bash
go run cmd/main.go
```

Server berjalan pada `http://localhost:8080`

## API Endpoints

| Method | Endpoint | Deskripsi |
|---|---|---|
| POST | `/article` | Membuat artikel baru |
| GET | `/article/:limit/:offset` | List artikel dengan paging |
| GET | `/article/:id` | Detail artikel |
| PUT | `/article/:id` | Update artikel |
| DELETE | `/article/:id` | Hapus artikel |

## API Testing

Import file `SharingVision.postman_collection.json` ke Postman untuk mencoba seluruh endpoint.

## Demo

- **Live API**: https://sharing-vision-backend-production.up.railway.app
- **Frontend repo**: https://github.com/Awanda2124/sharing-vision-frontend