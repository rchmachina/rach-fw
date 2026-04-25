# 🚀 Go Backend Framework

A backend service built with **Golang** using a **Clean Architecture** approach. The REST API layer is powered by **Gin**, providing a fast and minimal HTTP framework. This project is designed to be scalable, maintainable, and production-ready.

---

# 🧠 Overview

This project separates responsibilities across multiple layers:

* **Delivery Layer** → handles incoming requests (REST / gRPC)
* **Usecase Layer** → contains business logic
* **Repository Layer** → handles database operations
* **DTO Layer** → defines contracts

This structure helps the codebase become:

* Clean and organized
* Easy to scale
* Easier to test and maintain

---

# 📦 Project Structure

## 📁 `cmd/`

Application entry point

```
cmd/
 ├── rest/     → runs REST API server
 └── grpc/     → (optional) gRPC server
```

---

## 📁 `internal/`

Contains core application logic

### 📁 `delivery/`

Handles all incoming requests

```"
delivery/
 ├── rest/
 │    ├── handler/     → endpoint logic
 │    ├── middleware/  → auth, logging, etc.
 │    ├── routes/      → API route definitions
 │    └── router.go
 └── grpc/             → gRPC handlers (cooming soon)
```

---

### 📁 `usecase/`

Main business logic layer

```id="c3l8n1"
usecase/
 └── service.go
```

All business rules should live here, not in rest/grpc handler so all delivery method can use this usecase.

---

### 📁 `repository/`

Database access layer

```"
repository/
 ├── users/
 ├── ai/
 └── master_schema.go
```

This layer is responsible for handling all database interactions, including querying and data persistence. It ensures that the application logic remains decoupled from the underlying data source.

Each folder (e.g., users, ai) represents a specific domain or module, containing its own set of queries and data access logic.

master_schema.go is responsible for managing and separating database schemas, allowing the application to organize and isolate data across different domains or modules (e.g., multiple schemas within a single database or you can customize using two or more datbase).

---

### 📁 `dto/`

Data Transfer Objects

```
dto/
 ├── validation/
 │      ├──rest → API request structures validation
 └── model/     → data models
```

---

### 📁 `app/`

Dependency Injection container

```
app/
 └── container.go
```

Responsible for initializing and wiring dependencies.

---

### 📁 `utils/`

Helpers and utilities

```id="g7w9x1"
utils/
 ├── jwt/       → JWT utilities
 └── helper/    → general helpers
```

---

## 📁 `configs/`

Application configuration

```id="h8y2z3"
configs/
 ├── db.go
 ├── redis.go
 └── config_master.go
```

---

# ⚙️ Getting Started

## 1. Clone Repository

```
git clone https://github.com/rchmachina/rach-fw
cd  github.com/rchmachina/rach-fw
```

---

## 2. Setup Environment

Edit the `.env` file located at:

```id="j1c6d7"
cmd/rest/.env
```

Configure:

* Database connection
* Application port
* JWT secret
* Redis (optional)

---

## 3. Install Dependencies

```bash id="k2e8f9"
go mod tidy
```

---

## 4. Run Application

```bash"
go run cmd/rest/rest.go
```

Or use hot reload:

```bash"
air
```

---

# 🔐 Authentication

This project includes:

* JWT Authentication
* Middleware protection

Locations:

```id="n5k4l6"
internal/utils/jwt/
internal/delivery/rest/middleware/
```

---

# 🔄 Request Flow

Request lifecycle:

```id="o6m7n8"
Client → Handler → Usecase → Repository → Database
```

Example:

```id="p7o9q0"
POST /users
→ handler
→ usecase
→ repository
→ database
```

---

# 🧪 Adding a New Feature

Steps to add a new endpoint:

### 1. Create DTO

```id="q8r1s2"
dto/request/
```

### 2. Add Handler

```id="r9t3u4"
delivery/rest/handler/
```

### 3. Implement Usecase

```id="s0v5w6"
usecase/
```

### 4. Add Repository (if needed)

```id="t1x7y8"
repository/
```

### 5. Register Route

```id="u2z9a0"
delivery/rest/routes/
```

---

# 📌 Notes

* Ready for scaling into microservices
* Easy to extend with gRPC
* Clear separation of concerns

---

# 🚀 Coming soon Future 

* Unit Testing
* Swagger / OpenAPI documentation
* Docker setup
* gRPC implementation
* CI/CD pipeline

---

# 🤝 Contributing

Contributor are welcome 🚀
