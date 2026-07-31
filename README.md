# Restaurant Management System

A simple full-stack Restaurant Management System with a Go backend and a React + Vite frontend. The backend provides REST controllers for users, products, orders, payments, suppliers, inventory, purchases, and customers; the frontend is a Vite-powered React app that talks to the backend API.

---

## Table of contents
- [Features](#features)
- [Stack](#stack)
- [Repository structure](#repository-structure)
- [How it fits together](#how-it-fits-together)
- [Quickstart — run locally](#quickstart--run-locally)
  - [Backend (Go)](#backend-go)
  - [Frontend (React + Vite)](#frontend-react--vite)
- [Environment variables (example)](#environment-variables-example)
- [API overview](#api-overview)
- [Migrations](#migrations)
- [Development notes](#development-notes)
- [Contributing](#contributing)
- [License & contact](#license--contact)

---

## Features
- REST API for core restaurant operations:
  - Users, authentication (controller present)
  - Products & inventory management
  - Orders and order details
  - Payments and purchase flows
  - Suppliers and customer management
- Database migrations integrated with the backend
- Vite + React frontend scaffolded for the client UI

---

## Stack
- Language(s): Go (backend), JavaScript/React (frontend)
- Frameworks / runtime:
  - Backend: Go HTTP server (net/http) with project-structured packages
  - Frontend: React + Vite
- Notable libraries (inferred from repo layout):
  - godotenv (env file loading)
  - Vite (frontend dev server/build)
  - ESLint for frontend linting

---

## Repository structure
Top-level layout (only the main pieces shown):

```
Backend/                # Go backend
  main.go               # entrypoint delegating to cmd.Serve
  cmd/
    server.go           # server bootstrap: load config, DB, migrations, register routes, start http server (port 8080)
  go.mod
  go.sum
  pkg/
    config/             # app configuration helpers
    connection/         # DB connection (GetDB)
    controllers/        # HTTP handlers: users, products, orders, payments, suppliers, inventory, purchase, customers
      customers.controller.go
      inventory.controller.go
      inventory.product.go
      order.controller.go
      order.deatils.controller.go
      payment.controller.go
      products.controller.go
      purchase.controller.go
      supplier.controller.go
      users.controller.go
    core/               # core app initialization and wiring
    middleware/         # HTTP middleware (auth, logging, CORS, etc.)
    migrations/         # DB migrations runner
    models/             # DB models / structs
    repositories/       # DB persistence layer
    routers/             # route registration (RootRoutes)
    services/           # business logic / service layer
    utils/               # helpers
Frontend/               # React + Vite frontend
  index.html
  package.json
  package-lock.json
  vite.config.js
  public/               # static assets
  src/                  # app source (React components)
  README.md             # frontend-specific notes
```

How it fits together:
- The backend bootstrap (Backend/main.go → Backend/cmd/server.go) loads environment configuration, establishes a DB connection, runs migrations, initializes controllers and routers, and starts an HTTP server (listening on port 8080).
- Controllers in pkg/controllers handle incoming HTTP requests and delegate to services/repositories to interact with the database.
- The frontend is a Vite React app that serves the client UI (development server via `npm run dev`) and will call the backend API endpoints.

---

## How to run it

Prerequisites:
- Go 1.18+ (or whatever the go.mod requires)
- Node 16+ / npm (for frontend)
- A running relational database (Postgres is typical — set DB env vars)
- .env file with the required environment variables (see example below)

Open two terminals (one for backend, one for frontend):

Backend (from repo root)
```bash
cd Backend
# download modules
go mod download

# run (development)
go run .

# build
go build -o restaurant-server .
./restaurant-server
```
The server bootstrap listens on port 8080 by default (server.go prints "Server Running on Port 8080").

Frontend
```bash
cd Frontend
npm install
npm run dev
```
The Vite dev server will start (default: http://localhost:5173). Open the URL printed by Vite and ensure the frontend is configured to call the backend API (CORS or proxy).

---

## Environment variables (example)
The backend uses environment variables (godotenv is used in server bootstrap). Create a `.env` in Backend/ with values like:

```
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=restaurant_db

# App
APP_PORT=8080
# any JWT_SECRET or other auth-related secrets the project expects
JWT_SECRET=replace-with-secret
```

Adjust names as needed to match the config package in Backend/pkg/config.

---

## API overview
The backend exposes REST controllers for the following domains (controller files in Backend/pkg/controllers):
- Users (users.controller.go)
- Products (products.controller.go)
- Inventory (inventory.controller.go and inventory.product.go)
- Orders and order details (order.controller.go, order.deatils.controller.go)
- Payments (payment.controller.go)
- Purchases (purchase.controller.go)
- Suppliers (supplier.controller.go)
- Customers (customers.controller.go)

Example (generic) curl:
```bash
# list products (replace base URL/path according to routers)
curl http://localhost:8080/products
# create order
curl -X POST http://localhost:8080/orders -H "Content-Type: application/json" -d '{"user_id":1,"items":[...]}'
```
Check the routers implementation (Backend/pkg/routers) for exact paths and API versioning (e.g., `/api/...`) before calling.

---

## Migrations
The server bootstrap calls the migrations runner (Backend/pkg/migrations). On startup, the app will attempt to run migrations against the configured DB. Ensure the DB env vars are set and the database is reachable before starting the backend.

---

## Development notes
- Entry points:
  - Backend: `Backend/main.go` (calls `cmd.Serve()` in `Backend/cmd/server.go`)
  - Frontend: `Frontend/index.html`, `Frontend/src/` for React components
- The backend follows a layered structure: controllers → services → repositories → models.
- Middleware (logging, auth, CORS) is centralized under `pkg/middleware`.
- Add a `.env.example` file with required environment variable names to help other developers get started.

-
