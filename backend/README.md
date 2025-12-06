# Backend API - Zguba.gov

Backend API w FastAPI dla aplikacji do zgłaszania znalezionych rzeczy.

## Instalacja

1. Utwórz wirtualne środowisko:
```bash
python -m venv venv
source venv/bin/activate  # Na Windows: venv\Scripts\activate
```

2. Zainstaluj zależności:
```bash
pip install -r requirements.txt
```

3. Skopiuj plik konfiguracyjny:
```bash
cp .env.example .env
```

4. Zainicjalizuj bazę danych:
```bash
python init_db.py
```

## Uruchomienie

```bash
bash start.sh
```

Alternatywnie:
```bash
uvicorn main:app --reload --host 0.0.0.0 --port 8000
```

API będzie dostępne pod adresem: http://localhost:8000
Dokumentacja API: http://localhost:8000/docs

## Struktura projektu

```
backend/
├── main.py                 # Główny plik aplikacji
├── config.py              # Konfiguracja
├── database.py            # Konfiguracja bazy danych
├── init_db.py             # Inicjalizacja bazy danych
├── start.sh               # Skrypt startowy
├── models/                # Modele SQLAlchemy
│   └── found_item.py
├── schemas/               # Modele Pydantic
│   └── found_item.py
└── routers/               # Endpointy API
    ├── found_items.py
    └── stats.py
```

## API Endpoints

### Znalezione rzeczy
- `GET /api/found-items` - Lista znalezionych rzeczy
- `POST /api/found-items` - Dodaj nową rzecz
- `GET /api/found-items/{id}` - Szczegóły rzeczy
- `PUT /api/found-items/{id}` - Aktualizuj rzecz
- `DELETE /api/found-items/{id}` - Usuń rzecz
- `GET /api/found-items/categories/list` - Lista kategorii

### Statystyki
- `GET /api/stats` - Ogólne statystyki systemu

### Health Check
- `GET /health` - Status API

## Uwaga

Jednostki terytorialne są obsługiwane bezpośrednio przez frontend (z pliku JSON).
Backend zajmuje się tylko zarządzaniem znalezionymi rzeczami i statystykami.
