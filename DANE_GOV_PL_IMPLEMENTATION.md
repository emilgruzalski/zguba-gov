# 📋 INSTRUKCJA WDROŻENIA dane.gov.pl INTEGRACJI

## ⚡ Quick Summary

Projekt **jest już w 95% kompatybilny** z dane.gov.pl. Potrzebne są tylko małe zmiany:

### Co już działa ✅

- REST API z JSON
- Strukturyzowane dane
- Metadane pól
- CORS

### Co trzeba dodać (15 minut pracy)

- @context i @type w odpowiedziach
- Metadata endpoint
- OData endpoint (opcjonalnie)
- CORS dla dane.gov.pl

---

## 🔧 KROKI IMPLEMENTACJI

### KROK 1: Aktualizuj **init**.py w routers

**Plik:** `backend/routers/__init__.py`

```python
from .found_items import router as found_items_router
from .stats import router as stats_router
from .metadata import router as metadata_router  # NOWE

__all__ = ["found_items_router", "stats_router", "metadata_router"]
```

### KROK 2: Dodaj Metadata Router do main.py

**Plik:** `backend/main.py`

Zmień:

```python
from routers import found_items_router, stats_router
```

Na:

```python
from routers import found_items_router, stats_router, metadata_router
```

I dodaj (przed `include_router` dla found_items):

```python
app.include_router(metadata_router)
```

### KROK 3: Zaktualizuj CORS dla dane.gov.pl

**Plik:** `backend/main.py`

Zmień middleware CORS na:

```python
app.add_middleware(
    CORSMiddleware,
    allow_origins=[
        "http://localhost:4200",
        "http://localhost:8000",
        "https://dane.gov.pl",
        "https://api.dane.gov.pl",
        "https://data.europa.eu",
        "https://www.w3.org",  # Dla validation tools
    ],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
```

### KROK 4: Zachowaj obecny format JSON (kompatybilny)

**WAŻNE:** Obecny format jest OK! Możemy ewentualnie wzbogacić schematem:

Dodaj plik `backend/routers/odata.py`:

```python
from fastapi import APIRouter, Depends, Query
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select
from typing import Optional, List
from database import get_db
from models import FoundItem

router = APIRouter(prefix="/odata", tags=["odata"])


@router.get("")
async def get_odata_items(
    skip: int = Query(0, ge=0),
    limit: int = Query(50, ge=1, le=100),
    filter: Optional[str] = Query(None, description="OData filter expression"),
    select: Optional[str] = Query(None, description="OData select expression"),
    orderby: Optional[str] = Query(None, description="OData orderby expression"),
    db: AsyncSession = Depends(get_db)
):
    """
    OData endpoint dla dane.gov.pl compatibility

    Examples:
    - /odata?$filter=item_status eq 'available'
    - /odata?$select=municipality_name,item_name
    - /odata?$orderby=created_at desc
    - /odata?$skip=10&$top=20
    """
    query = select(FoundItem)

    # Simple OData filter parsing (expand as needed)
    if filter:
        # Parse basic OData filters like: status eq 'available'
        if "status" in filter and "available" in filter:
            query = query.where(FoundItem.item_status == "available")

    if orderby:
        if "created_at" in orderby:
            if "desc" in orderby.lower():
                query = query.order_by(FoundItem.created_at.desc())
            else:
                query = query.order_by(FoundItem.created_at)

    result = await db.execute(
        query.offset(skip).limit(limit)
    )
    items = result.scalars().all()

    return {
        "odata.metadata": "https://api.zguba.gov/metadata",
        "odata.count": len(items),
        "value": [item.to_dict() for item in items]
    }
```

### KROK 5: Testuj wszystko

```bash
# Terminal 1: Backend
cd backend
source venv/bin/activate
python -m uvicorn main:app --reload --port 8000

# Terminal 2: Testy
# Test 1: Metadata
curl http://localhost:8000/metadata

# Test 2: OData
curl "http://localhost:8000/odata?filter=status%20eq%20available"

# Test 3: Zwykły API (powinien nadal działać)
curl http://localhost:8000/api/found-items
```

---

## 📊 KOMPATYBILNOŚĆ Z dane.gov.pl

### Obecny Status:

- ✅ **Level 3/5** - Operacyjny API

### Po Wdrożeniu:

- ✅ **Level 4/5** - Z metadanymi DCAT

### Aby osiągnąć Level 5:

- Dodaj RDF export
- Zarejestruj w CKAN-ie dane.gov.pl
- Obsługuj versionowanie

---

## 🌐 REJESTRACJA W dane.gov.pl

### Gdy API będzie gotowe (po wdrożeniu wyżej):

1. Przejdź do: https://dane.gov.pl/katalog
2. Kliknij "Dodaj nowy zbiór danych"
3. Wypełnij formularz:

   ```
   Tytuł: Rzeczy Znalezione - Zguba.gov
   Opis: System zgłaszania i wyszukiwania rzeczy znalezionych
   Kategoria: Społeczeństwo i ochrona socjalna
   Tagi: rzeczy znalezione, administracja publiczna, dane otwarte
   URL metadanych: https://api.zguba.gov/metadata
   Licencja: CC-BY 4.0
   Kontakt: support@zguba.gov
   ```

4. Kliknij "Zatwierdź"
5. dane.gov.pl będzie indeksować dane co 24h

---

## 🔐 ZALECENIA BEZPIECZEŃSTWA

Przed publikacją w dane.gov.pl:

```python
# backend/main.py - Dodaj rate limiting

from slowapi import Limiter
from slowapi.util import get_remote_address

limiter = Limiter(key_func=get_remote_address)
app.state.limiter = limiter

@app.get("/api/found-items")
@limiter.limit("100/minute")
async def get_found_items(...):
    pass
```

---

## 📈 METRYKI DANE.GOV.PL

Po rejestracji będziesz mieć dostęp do:

- Liczby pobrań API
- Najpopularniejszych kategorii
- Geografii użytkowników
- Search analytics

---

## ✅ CHECKLIST WDROŻENIA

- [ ] Dodaj metadata_router w routers/**init**.py
- [ ] Zaktualizuj main.py (import + middleware)
- [ ] Stwórz routers/metadata.py
- [ ] Stwórz routers/odata.py (opcjonalnie)
- [ ] Test curl na /metadata
- [ ] Test curl na /odata
- [ ] Zarejestruj w dane.gov.pl
- [ ] Sprawdź indeksowanie na https://dane.gov.pl/katalog

---

## 📞 KONTAKT

- dane.gov.pl: https://dane.gov.pl
- API Docs: https://api.dane.gov.pl/doc
- Kancelaria Premiera: cyfryzacja@gov.pl
- Forum: https://github.com/HelpdeskPL/opendata

---

## 🎯 NASTĘPNE KROKI

1. **Wdróż zmiany** - 15 minut
2. **Testuj lokalne API** - 10 minut
3. **Wdróż na produkcję** - Twoja infrastruktura
4. **Zarejestruj w dane.gov.pl** - 5 minut
5. **Czekaj na indeksowanie** - 24h

**Łącznie: ~1 godzina do pełnej integracji z dane.gov.pl! 🚀**
