# Farmako - Coupon System

Farmako is a backend system for managing and validating coupons in a medicine ordering platform. This project is built in Go and follows a modular, production-ready architecture.

## Architecture

The architecture is designed to ensure modularity, scalability, and maintainability:
```
┌─────────────────┐
│   Client (API)  │
└────────┬────────┘
         │
┌────────▼────────┐
│  Gin Framework  │
└────────┬────────┘
         │
┌────────▼────────┐
│  Service Layer  │
└────────┬────────┘
         │
┌────────▼────────┐
│Repository Layer │
└────────┬────────┘
         │
┌────────▼────────┐
│PostgreSQL+Cache │
└─────────────────┘
```
## Folder Structure
```
Farmako/
├── apploader/         # Application initialization
├── cache/             # Cache management
├── controller/        # API routes and controllers
├── database/          # Database connection and repositories
│   └── migration/     # Database migration files
├── docs/              # Swagger documentation
├── handler/           # Request handlers
├── models/            # Data models
├── service/           # Business logic
├── utils/             # Utility functions
├── config/            # Configuration files
├── Dockerfile         # Container configuration
├── go.mod             # Go module dependencies
├── go.sum             # Go module checksums
└── main.go            # Application entry point
```
## Setup Instructions

### Running with Docker

1. Start the PostgreSQL database:
```
docker run --name coupon-store -e POSTGRES_PASSWORD=root -e POSTGRES_DB=coupon-service -p 5432:5432 -d postgres:latest
```
3. Run database migrations:
```
migrate -database "postgres://postgres:root@localhost:5432/coupon-service?sslmode=disable" -path ./database/migration up
```
5. Build and start the application:
```
docker build -t farmako-app .
docker run -d -p 8080:8080 --name farmako-app farmako-app
```
6. Access the application:
   - Swagger UI: ```http://localhost:8080/swagger/index.html```
   - API Endpoints: ```http://localhost:8080```


## API Endpoints

### 1. Create a Coupon
POST /coupons
```
Request Body:
{
    "coupon_code": "SAVE20",
    "expiry_date": "2025-05-19 18:54:15",
    "applicable_medicine_ids":["med_123","random"],
    "applicable_categories":["painkillers", "painrevivers"],
    "min_order_value":20,
    "terms_and_conditions":"T&C applies",
    "discount_type":"flat",
    "discount_value":20,
    "max_usage_per_user":2
}
```
```
Response:
{
    "message": "Coupon created successfully"
}
```
### 2. Get Applicable Coupons
GET /coupons/applicable
```
Request Body:
{
    "cart_items": [
        { "id": "med_123", "category": "painkiller" }
    ],
    "order_total": 700
}
```
```
Response:
{
    "applicable_coupons": [
        {
            "coupon_code": "SAVE20",
            "discount_value": 20
        }
    ]
}
```
### 3. Validate a Coupon
POST /coupons/validate
```
Request Body:
{
    "coupon_code": "SAVE20",
    "cart_items": [
        { "id": "med_123", "category": "painkiller" }
    ],
    "order_total": 700,
    "timestamp": "2025-05-05T15:00:00Z"
}
```
```
Success Response:
{
    "is_valid": true,
    "discount": {
        "items_discount": 140,
        "charges_discount": 0
    },
    "message": "Coupon applied successfully"
}
```
```
Failure Response:
{
  "is_valid": false,
  "reason": "Invalid coupon code"
}
```
### 4. Fetch All Cached Coupons
GET /coupon/cache
```
Response:
{
    "cache": {
        "SAVE20": {
            "coupon_code": "SAVE20",
            "expiry_date": "2025-05-19 18:54:15",
            "usage_type": "one_time",
            "applicable_medicine_ids": ["mde","bjr"],
            "applicable_categories": ["medics", "ores"],
            "min_order_value": 20,
            "valid_time_window": "24h",
            "terms_and_conditions": "Cannot be used on medical devices",
            "discount_type": "percentage",
            "discount_value": 20,
            "max_usage_per_user": 2
        }
    }
}
```
## Key Features

- Coupon Creation: Create and manage discount coupons
- Validation: Validate coupons against order criteria
- Caching: In-memory caching for fast coupon lookups
- Concurrency: Thread-safe operations with mutex locks
- Persistence: PostgreSQL database for reliable storage
- API Documentation: Swagger UI for interactive API exploration

## Technologies Used

- Go: Backend programming language
- Gin: Web framework
- PostgreSQL: Database
- Swagger: API documentation
- Docker: Containerization
- Migrate: Database migration tool

## Design Patterns

- Repository Pattern: Abstracts data access logic
- Service Layer: Contains business logic
- Dependency Injection: Promotes loose coupling
- Middleware Pattern: For request processing pipeline

## Concurrency & Caching

- RWMutex: Used for thread-safe cache operations
- In-Memory Cache: Fast access to frequently used coupon data
- Transaction Safety: Ensures data consistency during concurrent operations