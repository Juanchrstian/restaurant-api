# 🍽️ Restaurant API

A backend REST API for an online restaurant ordering system built with **Go (Golang)**. This project is designed using **Clean Architecture**, **Dependency Injection**, and **Feature-Based Structure** while demonstrating production-ready backend concepts such as Redis caching, database transactions, race condition handling, and idempotency.

---

## ✨ Features

### Current

- ✅ REST API using Gin
- ✅ PostgreSQL
- ✅ Redis
- ✅ Docker Compose
- ✅ Health Check API
- ✅ Environment Configuration (.env)
- ✅ Clean Project Structure

### Planned

- ⏳ Menu Management
- ⏳ Session Management (Table)
- ⏳ Online Ordering
- ⏳ Menu Add-ons
- ⏳ Redis Cache
- ⏳ Database Transaction
- ⏳ Race Condition Handling
- ⏳ Idempotency
- ⏳ API Documentation (Swagger)

---

# 🏗️ Architecture

```
Client
   │
   ▼
Gin Router
   │
   ▼
Handler
   │
   ▼
Service
   │
   ▼
Repository
   │
   ▼
PostgreSQL
```

Redis is used as a caching layer.

```
Client
   │
   ▼
Redis
   │
 Cache Hit?
   │
 ┌─Yes──────────────┐
 │                  ▼
 │          Return Response
 │
 └─No───────────────► PostgreSQL
                         │
                         ▼
                  Save to Redis
                         │
                         ▼
                  Return Response
```

---

# 📂 Project Structure

```
restaurant-api/

├── cmd/
│   └── api/
│       └── main.go
│
├── docker/
├── docs/
├── migrations/
├── seed/
│
├── internal/
│
│   ├── health/
│   ├── menu/
│   ├── order/
│   ├── session/
│
│   └── shared/
│       ├── config/
│       ├── constants/
│       ├── database/
│       ├── errors/
│       ├── logger/
│       ├── middleware/
│       ├── redis/
│       ├── response/
│       └── validator/
│
├── docker-compose.yml
├── go.mod
└── README.md
```

---

# 🚀 Tech Stack

| Category | Technology |
|----------|------------|
| Language | Go 1.24 |
| Framework | Gin |
| Database | PostgreSQL 17 |
| Cache | Redis 7 |
| ORM | GORM |
| Container | Docker |
| Configuration | godotenv |
| Database Migration | golang-migrate *(planned)* |

---

# 🚀 Getting Started

## Clone Repository

```bash
git clone https://github.com/juanchrstian/restaurant-api.git
cd restaurant-api
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Configure Environment

Copy

```
.env.example
```

to

```
.env
```

Example

```env
APP_NAME=Restaurant API
APP_ENV=development
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=restaurant
DB_PASSWORD=restaurant
DB_NAME=restaurant_db
DB_SSLMODE=disable

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

---

## Start PostgreSQL & Redis

```bash
docker compose up -d
```

---

## Run Application

```bash
go run ./cmd/api
```

---

# 📡 API

## Health Check

```
GET /api/v1/health
```

Response

```json
{
    "status": "UP"
}
```

---

# 📅 Development Roadmap

## Sprint 1

- [x] Docker
- [x] PostgreSQL
- [x] Redis
- [x] Config
- [x] Database Connection
- [x] Redis Connection
- [x] Router
- [x] Health API

## Sprint 2

- [ ] Database Migration
- [ ] Seeder
- [ ] Menu CRUD
- [ ] Redis Cache

## Sprint 3

- [ ] Table Session
- [ ] Customer Session
- [ ] Shared Order

## Sprint 4

- [ ] Order API
- [ ] Transaction
- [ ] Race Condition
- [ ] Idempotency

## Sprint 5

- [ ] Authentication
- [ ] Logging
- [ ] Validation
- [ ] Swagger

---

# 📖 Learning Goals

This project is built to explore production-ready backend development with Go, including:

- Clean Architecture
- Repository Pattern
- Dependency Injection
- REST API Design
- PostgreSQL
- Redis Caching
- Docker
- Database Transactions
- Concurrency
- Race Condition Handling
- Idempotency
- Backend Best Practices

---

# 👨‍💻 Author

**Juan Christian**

GitHub: https://github.com/juanchrstian
