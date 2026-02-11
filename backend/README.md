# Backend API - Zguba.gov

FastAPI backend for the lost & found reporting application.

## Setup

```bash
uv venv
source .venv/bin/activate  # Windows: .venv\Scripts\activate
uv pip install -r requirements.txt
cp .env.example .env
python init_db.py
```

## Running

```bash
uvicorn main:app --reload --host 0.0.0.0 --port 8000
```

Or with Docker:

```bash
docker build -t zguba-backend .
docker run -p 8000:8000 zguba-backend
```

- API: http://localhost:8000
- Swagger docs: http://localhost:8000/docs

## Project Structure

```
backend/
├── main.py              # Application entry point
├── config.py            # Settings (from .env)
├── database.py          # Database engine and session
├── init_db.py           # Database initialization
├── add_examples.py      # Seed example data
├── Dockerfile
├── models/
│   └── found_item.py    # SQLAlchemy ORM model
├── schemas/
│   ├── found_item.py    # Pydantic request/response schemas
│   └── found_item_dcat.py  # DCAT-AP schemas for dane.gov.pl
└── routers/
    ├── found_items.py   # CRUD endpoints
    ├── stats.py         # Statistics endpoint
    ├── odata.py         # OData-compatible endpoint
    └── metadata.py      # DCAT-AP metadata endpoint
```

## API Endpoints

### Found Items
- `GET /api/found-items` - List items (supports `category`, `municipality`, `status`, `search` filters)
- `POST /api/found-items` - Create item
- `GET /api/found-items/{id}` - Get item
- `PUT /api/found-items/{id}` - Update item
- `DELETE /api/found-items/{id}` - Delete item
- `GET /api/found-items/categories/list` - List distinct categories

### Statistics
- `GET /api/stats` - Item counts, top categories, top municipalities

### OData
- `GET /odata/FoundItems` - OData query interface (`$filter`, `$orderby`, `$top`, `$skip`, `$count`)

### Metadata
- `GET /metadata` - DCAT-AP catalog metadata
- `GET /metadata/schema` - JSON schema for found items
- `GET /metadata/dcat` - RDF endpoint (placeholder)
- `GET /metadata/distribution/{id}` - Distribution metadata

### Health
- `GET /health` - Returns `{"status": "ok"}`

## Note

Territorial units are handled by the frontend (from a bundled JSON file). The backend manages found items and statistics only.
