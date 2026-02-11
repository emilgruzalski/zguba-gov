# Zguba.gov - Lost & Found Reporting System

A web application for reporting and searching found items in Polish public administration offices.

Built in 24 hours during the [HackNation](https://hacknation.pl/#about) hackathon.

## Project Structure

```
zguba-gov/
├── cmd/server/         # Application entrypoint
├── internal/
│   ├── config/         # Environment configuration
│   ├── database/       # SQLite database setup
│   ├── handler/        # HTTP handlers (API, OData, metadata, pages)
│   ├── model/          # Data models
│   ├── municipality/   # Territorial unit service + data
│   ├── odata/          # OData query parser
│   └── repository/     # Database repository layer
├── web/
│   ├── static/         # CSS, images
│   └── templates/      # Go HTML templates (HTMX)
├── Dockerfile
├── docker-compose.yml
└── go.mod
```

## Requirements

- Go 1.24+

## Getting Started

### Option 1: Docker (recommended)

```bash
docker compose up --build
```

Application: http://localhost:8000

### Option 2: Local Development

```bash
go run ./cmd/server
```

Application: http://localhost:8000

### Environment Variables

| Variable       | Default            | Description          |
|----------------|--------------------|----------------------|
| `PORT`         | `8000`             | Server port          |
| `DATABASE_URL` | `zguba_gov.db`     | SQLite database path |
| `CORS_ORIGINS` | `http://localhost:4200,http://localhost:3000` | Allowed CORS origins |

## Features

- Multi-step wizard for registering found items (HTMX-powered)
- Territorial unit autocomplete (voivodeships, counties, municipalities)
- Items listing with filtering and search
- JSON/CSV export
- RESTful API with full CRUD
- OData-compatible endpoint with `$filter`, `$orderby`, `$top`, `$skip`, `$count`
- DCAT-AP metadata for dane.gov.pl integration
- Responsive UI following GOV.PL design guidelines
- Single binary with embedded static assets

## API Endpoints

### Found Items
- `GET /api/found-items` - List items (query: `skip`, `limit`, `category`, `municipality`, `status`, `search`)
- `POST /api/found-items` - Create item
- `GET /api/found-items/:id` - Get item
- `PUT /api/found-items/:id` - Update item
- `DELETE /api/found-items/:id` - Delete item
- `GET /api/found-items/categories/list` - List categories
- `GET /api/stats` - Statistics

### OData
- `GET /odata/FoundItems` - Query (`$filter`, `$orderby`, `$top`, `$skip`, `$count`)
- `GET /odata/$metadata` - EDMX metadata

### Metadata (DCAT-AP)
- `GET /metadata` - Dataset catalog
- `GET /metadata/distribution/:id` - Distribution info

### Health
- `GET /health` - Health check

## Tech Stack

- **Go 1.24** + **Gin** (HTTP framework)
- **SQLite** via `modernc.org/sqlite` (pure Go, no CGO)
- **Go HTML templates** + **HTMX** (server-rendered UI)
- **embed** (static assets bundled into binary)

## Development

### Adding a new API endpoint

1. Define model structs in `internal/model/`
2. Add repository methods in `internal/repository/`
3. Create handler methods in `internal/handler/`
4. Register routes in `cmd/server/main.go`

### Adding a new page/wizard step

1. Create a template in `web/templates/partials/`
2. Add handler logic in `internal/handler/pages.go`
3. Register routes in `cmd/server/main.go`

## Testing

```bash
go test ./...
```

## CI

GitHub Actions runs build, tests, and `go vet` on every push and PR to `main`.

## License

Apache License 2.0 - see [LICENSE](LICENSE) for details.
