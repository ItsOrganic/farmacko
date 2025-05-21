# Cloud-Sek Blog Application

A modern blog application with a Go backend using Gin framework, PostgreSQL database, and Markdown support.

## Architecture Overview

```
┌───────────┐     ┌───────────┐     ┌───────────┐     ┌───────────┐
│ HTTP API  │────▶│  Service  │────▶│ Repository│────▶│ PostgreSQL│
│ (Gin)     │◀────│  Layer    │◀────│  Layer    │◀────│  Database │
└───────────┘     └───────────┘     └───────────┘     └───────────┘
       │                 │
       │                 │
       ▼                 ▼
┌───────────┐     ┌───────────┐
│  Swagger  │     │ In-Memory │
│    Docs   │     │   Cache   │
└───────────┘     └───────────┘
Gin (HTTP API) <-> Service Layer <-> Repository <-> PostgreSQL
```
**Key Features:**
- Markdown support in posts/comments (converted to HTML for rich text display)
- Swagger API documentation
- Layered architecture for separation of concerns
- In-memory caching for improved performance

## Directory Structure

```
├── apploader/     # Configuration and cache loader
├── cache/         # In-memory cache implementation
├── config/        # YAML configuration for database
├── constants/     # SQL queries, HTML templates
├── controller/    # Gin route definitions
├── database/      # Database connection and repository logic
├── docs/          # Swagger documentation (auto-generated)
├── globals/       # Global configuration, database, and cache
├── handler/       # HTTP request handlers
├── models/        # Data models
├── service/       # Business logic layer
├── utils/         # Utilities (Markdown-to-HTML conversion)
├── main.go        # Application entrypoint
├── Dockerfile     # Container build definition
└── README.md      # This documentation
```

## Getting Started

### 1. Start PostgreSQL with Docker

```bash
docker run --name cloudsek-post-store -e POSTGRES_PASSWORD=root -e POSTGRES_DB=post-comments-service -p 5432:5432 -d postgres:latest
```

### 2. Run Database Migrations

```bash
# Install the migrate tool
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run migrations
migrate -database "postgres://postgres:root@localhost:5432/post-comments-service?sslmode=disable" -path ./database/migration up
```

### 3. Configure Database Connection

Edit if running local POSTGRES `config/config.yaml`:

```yaml
db:
  user-name: "postgres"
  password: "root"
  host: "localhost"
  port: 5432
  database: "post-comments-service"
  driver: "postgres"
  ssl-mode: "disable"
```

### 4. Build and Run Locally

```bash
go build -o blog-app .
./blog-app
```

### 5. Build and Run with Docker

```bash
# Build the Docker image
docker build -t cloudsek-app .

# Run the container
docker run -d -p 8080:8080 --name cloudsek-app cloudsek-app
```

## Usage

### Access Points

- **API**: http://localhost:8080
- **Swagger Documentation**: http://localhost:8080/swagger/index.html

### Example API Usage

#### Create a Post

```bash
curl -X POST http://localhost:8080/create \
  -H "Content-Type: application/json" \
  -d '{"title":"My First Post","description":"This is **bold** and *italic* text with a link"}'
```

#### Add a Comment

```bash
curl -X POST http://localhost:8080/post/1/comment \
  -H "Content-Type: application/json" \
  -d '{"author":"John","message":"This is a **great** post!"}'
```

#### View a Post

```bash
curl http://localhost:8080/post/1
```

#### View Comments (HTML)

```bash
curl http://localhost:8080/post/1/comments
```

## Troubleshooting

| Issue | Solution |
|-------|----------|
| Swagger error | Run `swag init` in project root |
| Database connection error | Check `config/config.yaml` and PostgreSQL container status |
| Docker container issues | Check logs with `docker logs cloudsek-app` |
| Missing Go dependencies | Run `go mod tidy` to install required packages |
