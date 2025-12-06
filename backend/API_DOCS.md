# API Documentation

## Base URL
```
http://localhost:8000
```

## Health Check

### GET /health
Sprawdza status API.

**Response:**
```json
{
  "status": "healthy"
}
```

## Znalezione Rzeczy (Found Items)

### GET /api/found-items
Pobiera listę znalezionych rzeczy.

**Query Parameters:**
- `skip` (int, optional): Liczba pominiętych rekordów (default: 0)
- `limit` (int, optional): Maksymalna liczba wyników (default: 50, max: 100)
- `category` (string, optional): Kategoria przedmiotu
- `municipality` (string, optional): Nazwa gminy
- `status` (string, optional): Status (available, claimed, expired)
- `search` (string, optional): Wyszukaj w nazwie lub opisie

**Example:**
```bash
curl "http://localhost:8000/api/found-items?category=dokumenty&status=available"
```

**Response:**
```json
[
  {
    "id": "uuid-here",
    "municipality": {
      "name": "Warszawa",
      "type": "miasto",
      "contactEmail": "kontakt@um.warszawa.pl"
    },
    "item": {
      "name": "Portfel",
      "category": "dokumenty",
      "date": "2025-12-01",
      "location": "Park Łazienkowski",
      "status": "available",
      "description": "Brązowy portfel ze skóry"
    },
    "pickup": {
      "deadline": 30,
      "location": "Urząd Miasta, ul. Senatorska 30",
      "hours": "9:00-17:00",
      "contact": "+48 22 443 44 44"
    },
    "categories": ["dokumenty", "portfele"],
    "createdAt": "2025-12-01T10:00:00",
    "updatedAt": "2025-12-01T10:00:00"
  }
]
```

### POST /api/found-items
Dodaje nową znalezioną rzecz.

**Request Body:**
```json
{
  "municipality": {
    "name": "Warszawa",
    "type": "miasto",
    "contactEmail": "kontakt@um.warszawa.pl"
  },
  "item": {
    "name": "Portfel",
    "category": "dokumenty",
    "date": "2025-12-01",
    "location": "Park Łazienkowski",
    "status": "available",
    "description": "Brązowy portfel ze skóry"
  },
  "pickup": {
    "deadline": 30,
    "location": "Urząd Miasta, ul. Senatorska 30",
    "hours": "9:00-17:00",
    "contact": "+48 22 443 44 44"
  },
  "categories": ["dokumenty", "portfele"]
}
```

**Example:**
```bash
curl -X POST "http://localhost:8000/api/found-items" \
  -H "Content-Type: application/json" \
  -d @example_item.json
```

**Response:** Status 201 + utworzony obiekt

### GET /api/found-items/{item_id}
Pobiera szczegóły znalezionej rzeczy.

**Example:**
```bash
curl "http://localhost:8000/api/found-items/uuid-here"
```

### PUT /api/found-items/{item_id}
Aktualizuje znalezioną rzecz.

**Request Body:** Taki sam jak POST (wszystkie pola opcjonalne)

**Example:**
```bash
curl -X PUT "http://localhost:8000/api/found-items/uuid-here" \
  -H "Content-Type: application/json" \
  -d '{"item": {"status": "claimed"}}'
```

### DELETE /api/found-items/{item_id}
Usuwa znalezioną rzecz.

**Example:**
```bash
curl -X DELETE "http://localhost:8000/api/found-items/uuid-here"
```

**Response:** Status 204 (No Content)

### GET /api/found-items/categories/list
Pobiera listę wszystkich używanych kategorii.

**Example:**
```bash
curl "http://localhost:8000/api/found-items/categories/list"
```

## Statystyki (Stats)

### GET /api/stats
Pobiera ogólne statystyki systemu.

**Example:**
```bash
curl "http://localhost:8000/api/stats"
```

**Response:**
```json
{
  "foundItems": {
    "total": 150,
    "available": 100,
    "claimed": 50
  },
  "topCategories": [
    {"category": "dokumenty", "count": 50},
    {"category": "telefony", "count": 30}
  ],
  "topMunicipalities": [
    {"name": "Warszawa", "count": 45},
    {"name": "Kraków", "count": 30}
  ]
}
```

## Uwaga

**Jednostki terytorialne** są obsługiwane bezpośrednio przez frontend z lokalnego pliku JSON.
Backend zajmuje się tylko zarządzaniem znalezionymi rzeczami i statystykami.

## Kody błędów

- `200` - OK
- `201` - Created
- `204` - No Content
- `400` - Bad Request (nieprawidłowe dane wejściowe)
- `404` - Not Found (zasób nie istnieje)
- `422` - Unprocessable Entity (błąd walidacji)
- `500` - Internal Server Error

## Interaktywna dokumentacja

FastAPI automatycznie generuje interaktywną dokumentację:

- **Swagger UI**: http://localhost:8000/docs
- **ReDoc**: http://localhost:8000/redoc
