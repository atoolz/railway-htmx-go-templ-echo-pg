<p align="center">
  <img src="assets/banner.png" alt="HTMX + Go + PostgreSQL Railway Template" width="100%">
</p>

<p align="center">
  <strong>Production-ready HTMX starter with Go, Templ, Echo, and PostgreSQL. One-click deploy to Railway.</strong>
</p>

<p align="center">
  <a href="https://railway.com/deploy/htmx-go-echo-pg">
    <img src="https://railway.com/button.svg" alt="Deploy on Railway">
  </a>
</p>

<p align="center">
  <a href="https://github.com/atoolz/railway-htmx-go-templ-echo-pg/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/atoolz/railway-htmx-go-templ-echo-pg?style=flat-square&color=00c9a7" alt="License">
  </a>
  <img src="https://img.shields.io/badge/Go-1.25-00ADD8?style=flat-square" alt="Go 1.25">
  <img src="https://img.shields.io/badge/HTMX-2.0.7-3366CC?style=flat-square" alt="HTMX 2.0.7">
  <img src="https://img.shields.io/badge/Echo-v4.15-00c9a7?style=flat-square" alt="Echo v4.15">
  <img src="https://img.shields.io/badge/PostgreSQL-16-336791?style=flat-square" alt="PostgreSQL">
</p>

<br>

## What's Inside

A complete, working Todo application that demonstrates the full HTMX + Go stack. No JavaScript frameworks, no build steps for the frontend, no client-side state management. Just HTML over the wire.

| Layer | Technology | Role |
|-------|-----------|------|
| **Frontend** | HTMX 2.0.7 + Tailwind CSS | Hypermedia-driven interactions via CDN |
| **Templating** | Templ v0.3 | Type-safe Go templates with IDE support |
| **Routing** | Echo v4.15 | High-performance HTTP framework |
| **Database** | PostgreSQL + pgx v5.9 | Connection pooling, prepared statements |
| **Deploy** | Dockerfile (multi-stage) | 31s build, ~15MB final image |

<br>

## Why This Stack

**HTMX** lets you build dynamic UIs by returning HTML fragments from the server instead of JSON. No React, no Vue, no bundle. Your Go backend renders HTML, HTMX swaps it into the DOM. The browser does what browsers do best.

**Templ** gives you type-safe, composable HTML components written in Go. Compile-time checking catches broken templates before they reach production. IDE autocompletion works because components are just Go functions.

**Echo** is one of the fastest Go web frameworks with a clean API, middleware support, and zero allocation routing. It handles the HTTP layer so you focus on your application.

**PostgreSQL** needs no introduction. The pgx driver gives you connection pooling, prepared statements, and direct access to PostgreSQL-specific features without an ORM in the way.

<br>

## Project Structure

```
.
├── cmd/server/
│   └── main.go              # Server setup, routes, middleware
├── internal/
│   ├── database/
│   │   └── database.go      # pgx connection pool + auto-migration
│   ├── handlers/
│   │   └── handlers.go      # HTTP handlers (CRUD + health check)
│   └── models/
│       └── models.go        # Data structures
├── templates/
│   ├── layout.templ          # Base HTML with HTMX + Tailwind CDN
│   ├── home.templ            # Home page composition
│   └── components.templ      # HTMX-powered interactive components
├── Dockerfile                # Multi-stage build (golang -> alpine)
└── go.mod
```

<br>

## HTMX Patterns Demonstrated

The Todo app showcases core HTMX patterns you'll use in real projects:

- **`hx-post` + `hx-target` + `hx-swap="afterbegin"`** - Create items and prepend to list without page reload
- **`hx-patch` + `hx-target` (self)** - Toggle completion with in-place swap
- **`hx-delete` + `hx-swap="outerHTML"`** - Remove items with smooth DOM removal
- **`hx-on::after-request`** - Reset form after successful submission
- **Health check endpoint** - `/health` returns JSON status with DB ping

<br>

## Deploy to Railway

Click the button above or:

1. Fork this repo
2. Create a new project on [Railway](https://railway.com)
3. Add a **PostgreSQL** database
4. Add a service from your forked repo
5. Set the variable: `DATABASE_URL` = `${{Postgres.DATABASE_URL}}`
6. Railway auto-detects the Dockerfile and deploys

The app auto-migrates the database on startup. No manual SQL needed.

<br>

## Local Development

```bash
# Prerequisites: Go 1.25+, PostgreSQL running locally

# Install templ CLI
go install github.com/a-h/templ/cmd/templ@latest

# Clone and run
git clone https://github.com/atoolz/railway-htmx-go-templ-echo-pg.git
cd railway-htmx-go-templ-echo-pg

# Set your database URL
export DATABASE_URL="postgres://user:pass@localhost:5432/mydb?sslmode=disable"

# Generate templ files and run
templ generate
go run ./cmd/server
```

Open [http://localhost:8080](http://localhost:8080).

<br>

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `DATABASE_URL` | Yes | - | PostgreSQL connection string |
| `PORT` | No | `8080` | Server port (Railway sets this automatically) |

<br>

## Part of the HTMX Railway Collection

This is one of 15 HTMX starter templates covering different backend stacks, all following the same pattern and ready for Railway deployment:

| Stack | Status |
|-------|--------|
| **Go + Echo** | This repo |
| Go + Fiber | Coming soon |
| Go + Chi | Coming soon |
| Python + FastAPI | Coming soon |
| Python + Django | Coming soon |
| Java + Spring Boot (PG) | Coming soon |
| Java + Spring Boot (MySQL) | Coming soon |
| PHP + Laravel | Coming soon |
| Node + Express | Coming soon |
| Ruby + Rails 8 | Coming soon |
| .NET + Razor | Coming soon |
| Rust + Axum | Coming soon |
| Node + Hono | Coming soon |
| Bun + Elysia | Coming soon |
| Elixir + Phoenix | Coming soon |

<br>

## License

[MIT](LICENSE)

---

<p align="center">
  <sub>Built by <a href="https://github.com/atoolz">AToolZ</a> for the HTMX community</sub>
</p>
