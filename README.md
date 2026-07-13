# Sharing Vision Backend Technical Test

## Requirements

- Go 1.24+
- MySQL
- golang-migrate

## How to Run

### 1. Clone repository

```bash
git clone https://github.com/<username>/sharing-vision-backend.git
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

Copy:

```
.env.example
```

menjadi

```
.env
```

Lalu sesuaikan konfigurasi database.

### 5. Run migration

```bash
migrate -path migrations -database "<YOUR_DATABASE_DSN>" up
```

### 6. Run application

```bash
go run cmd/main.go
```

Server berjalan pada:

```
http://localhost:8080
```

## API Testing

Import file berikut ke Postman:

```
SharingVision.postman_collection.json
```