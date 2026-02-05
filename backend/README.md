# Ecommerce Backend (Go)

🔧 **Overview**

This is the backend service for the ecommerce sample app. It exposes simple REST endpoints for managing products and users.

## Requirements

- Go 1.20+ (or latest stable)
- A .env file with required environment variables (see below)

## Environment variables

Required variables (example values):

```
VERSION=0.1.0
SERVICE_NAME=ecommerce
HTTP_PORT=8080
JWT_SECRET_KEY=my-secret-key
```

Place these values in a `.env` file in the `backend/` directory.

## How to run

From the `backend/` folder:

```bash
# format files
gofmt -w .
# validate code
go vet ./...
# run the server
go run main.go
```

The server will start on the port defined by `HTTP_PORT`.

## Endpoints

- GET /products — list all products
- POST /products — create a product (JSON body)
- GET /products/{id} — get product by id
- PUT /products/{id} — update product by id (protected by JWT)
- DELETE /products/{id} — delete product by id (protected by JWT)
- POST /users/login — login to receive a JWT token

Example: Create product

```bash
curl -X POST http://localhost:8080/products \
  -H 'Content-Type: application/json' \
  -d '{"name":"Widget","price":9.99}'
```

Example: Login

```bash
curl -X POST http://localhost:8080/users/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"user@example.com","password":"pass"}'
```

## Notes & Recent fixes

- Renamed `GetCongfig()` to `GetConfig()` to fix a typo.
- Fixed `GetProductsByID` to return HTTP 200 when a product is found.

If you'd like, I can also:
- Add more GoDoc comments across all packages
- Generate an OpenAPI/Swagger spec for these endpoints
- Add unit tests for handlers

---
