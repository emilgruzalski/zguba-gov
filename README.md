# Zguba.gov - System zgłaszania znalezionych rzeczy

Aplikacja webowa do zgłaszania i wyszukiwania znalezionych rzeczy w jednostkach administracji publicznej w Polsce.

## Struktura projektu

```
zguba-gov/
├── backend/          # Backend API (Python FastAPI)
├── frontend/         # Frontend (Angular)
└── README.md
```

## Wymagania

### Backend
- Python 3.9+
- pip

### Frontend
- Node.js 18+
- npm

## Instalacja i uruchomienie

### Backend (FastAPI)

1. Przejdź do katalogu backend:
```bash
cd backend
```

2. Utwórz wirtualne środowisko:
```bash
python3 -m venv venv
source venv/bin/activate  # Na Windows: venv\Scripts\activate
```

3. Zainstaluj zależności:
```bash
pip install -r requirements.txt
```

4. Skopiuj plik konfiguracyjny:
```bash
cp .env.example .env
```

5. Zainicjalizuj bazę danych:
```bash
python init_db.py
```

6. Uruchom serwer:
```bash
uvicorn main:app --reload --host 0.0.0.0 --port 8000
```

Backend będzie dostępny pod adresem: http://localhost:8000
Dokumentacja API: http://localhost:8000/docs

### Frontend (Angular)

1. Przejdź do katalogu frontend:
```bash
cd frontend
```

2. Zainstaluj zależności:
```bash
npm install
```

3. Uruchom serwer deweloperski:
```bash
npm start
```

Frontend będzie dostępny pod adresem: http://localhost:4200

## API Endpoints

### Znalezione rzeczy
- `GET /api/found-items` - Lista znalezionych rzeczy (z filtrami)
- `POST /api/found-items` - Dodaj nową rzecz
- `GET /api/found-items/{id}` - Szczegóły rzeczy
- `PUT /api/found-items/{id}` - Aktualizuj rzecz
- `DELETE /api/found-items/{id}` - Usuń rzecz
- `GET /api/found-items/categories/list` - Lista kategorii

### Statystyki
- `GET /api/stats` - Ogólne statystyki systemu

**Uwaga:** Jednostki terytorialne są obsługiwane bezpośrednio przez frontend (z pliku JSON).

## Funkcjonalności

### Zaimplementowane
- ✅ Przeglądanie jednostek terytorialnych
- ✅ Wyszukiwanie jednostek po nazwie
- ✅ Filtrowanie po województwach
- ✅ Zarządzanie znalezionymi rzeczami (CRUD)
- ✅ Filtrowanie znalezionych rzeczy
- ✅ Statystyki systemu
- ✅ API RESTful
- ✅ Dokumentacja API (Swagger)

### Frontend
- ✅ Formularz zgłaszania znalezionych rzeczy
- ✅ Autouzupełnianie jednostek terytorialnych
- ✅ Wieloetapowy formularz
- ✅ Walidacja danych
- ✅ Responsywny interfejs

## Technologie

### Backend
- FastAPI - framework webowy
- SQLAlchemy - ORM
- SQLite - baza danych (dev)
- Pydantic - walidacja danych
- Uvicorn - serwer ASGI

### Frontend
- Angular 18
- TypeScript
- RxJS
- Angular Forms
- HttpClient

## Rozwój

### Dodanie nowej funkcjonalności do API

1. Utwórz nowy model w `backend/models/`
2. Dodaj schemat Pydantic w `backend/schemas/`
3. Utwórz router w `backend/routers/`
4. Dodaj router do `backend/main.py`

### Dodanie nowej funkcjonalności do frontendu

1. Utwórz nowy komponent: `ng generate component nazwa-komponentu`
2. Dodaj routing w `app.routes.ts`
3. Utwórz serwis: `ng generate service services/nazwa-serwisu`

## Testowanie

### Backend
```bash
cd backend
source venv/bin/activate
pytest
```

### Frontend
```bash
cd frontend
npm test
```

## Licencja

MIT

## Autorzy

Projekt Zguba.gov - System zgłaszania znalezionych rzeczy
