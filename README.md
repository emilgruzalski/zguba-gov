# Zguba.gov - Lost & Found Reporting System

A web application for reporting and searching found items in Polish public administration offices.

Built in 2 days during the [HackNation](https://hacknation.pl/#about) hackathon.

## Project Structure

```
zguba-gov/
├── backend/           # Backend API (Python FastAPI)
├── frontend/          # Frontend (Angular 19)
├── docker-compose.yml # Docker orchestration
└── README.md
```

## Requirements

### Backend
- Python 3.9+
- [uv](https://docs.astral.sh/uv/)

### Frontend
- [Bun](https://bun.sh/)

## Getting Started

### Option 1: Docker (recommended)

Run the entire stack with Docker Compose:

```bash
docker compose up --build
```

- Frontend: http://localhost
- Backend API: http://localhost:8000
- API docs: http://localhost:8000/docs

### Option 2: Local Development

#### Backend (FastAPI)

```bash
cd backend
uv venv
source .venv/bin/activate  # Windows: .venv\Scripts\activate
uv pip install -r requirements.txt
cp .env.example .env
python init_db.py
uvicorn main:app --reload --host 0.0.0.0 --port 8000
```

- API: http://localhost:8000
- Swagger docs: http://localhost:8000/docs

#### Frontend (Angular)

```bash
cd frontend
bun install
bun start
```

- App: http://localhost:4200

## API Endpoints

### Found Items
- `GET /api/found-items` - List found items (with filters)
- `POST /api/found-items` - Create a new item
- `GET /api/found-items/{id}` - Get item details
- `PUT /api/found-items/{id}` - Update an item
- `DELETE /api/found-items/{id}` - Delete an item
- `GET /api/found-items/categories/list` - List categories

### Statistics
- `GET /api/stats` - System statistics

### OData
- `GET /odata/FoundItems` - OData-compatible endpoint

### Metadata (DCAT-AP)
- `GET /metadata` - Dataset metadata in DCAT-AP format
- `GET /metadata/schema` - JSON schema
- `GET /metadata/dcat` - RDF metadata
- `GET /metadata/distribution/{id}` - Distribution metadata

### Health
- `GET /health` - API health check

**Note:** Territorial units are handled directly by the frontend (from a bundled JSON file).

## Features

### Backend
- RESTful API with full CRUD
- OData-compatible endpoint
- DCAT-AP metadata for dane.gov.pl integration
- Automatic API documentation (Swagger / ReDoc)
- Async SQLAlchemy with SQLite

### Frontend
- Multi-step form for reporting found items
- Territorial unit autocomplete (voivodeships, counties, municipalities)
- Item filtering and search
- JSON/CSV export
- Responsive UI following GOV.PL design guidelines

## Tech Stack

### Backend
- FastAPI
- SQLAlchemy (async)
- SQLite + aiosqlite
- Pydantic

### Frontend
- Angular 19
- TypeScript
- RxJS
- Angular Reactive Forms

## Development

### Adding a new API feature

1. Create a model in `backend/models/`
2. Add a Pydantic schema in `backend/schemas/`
3. Create a router in `backend/routers/`
4. Register the router in `backend/main.py`

### Adding a new frontend feature

1. Generate a component: `ng generate component component-name`
2. Add routing in `app.routes.ts`
3. Generate a service: `ng generate service services/service-name`

## Testing

```bash
# Backend
cd backend && pytest

# Frontend
cd frontend && bun test
```

## License

Apache License 2.0 - see [LICENSE](LICENSE) for details.
