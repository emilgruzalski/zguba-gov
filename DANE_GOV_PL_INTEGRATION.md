# Integracja z dane.gov.pl - Analiza Zgodności Formatu JSON

## 📊 Obecny Format JSON Projektu vs. Wymogi dane.gov.pl

### Status: ✅ ZGODNY (z rekomendacjami)

---

## 1. STANDARD dane.gov.pl

dane.gov.pl wymaga zgodności z:

- **OData** - otwarte dane publiczne
- **DCAT** (Data Catalog Vocabulary) - katalog zbiorów danych
- **Schemat JSON** - strukturyzowane dane

### Wymagane Metadane:

```json
{
  "@context": "https://www.w3.org/ns/dcat",
  "dct:title": "Rzeczy Znalezione",
  "dct:description": "System zgłaszania znalezionych przedmiotów",
  "dcat:dataset": [
    /* dane */
  ],
  "dct:issued": "2025-12-01T00:00:00Z",
  "dct:modified": "2025-12-01T00:00:00Z"
}
```

---

## 2. NASZ OBECNY FORMAT

### ✅ Zalety:

- ✅ Hierarchiczna struktura (municipality, item, pickup)
- ✅ Logiczny podział odpowiedzialności
- ✅ Pola obowiązkowe i opcjonalne
- ✅ ISO 8601 dla dat (createdAt, updatedAt)
- ✅ Walidacja typu email (contactEmail)

```json
{
  "id": "uuid",
  "municipality": {
    "name": "string",
    "type": "string",
    "contactEmail": "email@example.com"
  },
  "item": {
    "name": "string",
    "category": "string",
    "date": "YYYY-MM-DD",
    "location": "string",
    "status": "string",
    "description": "string (optional)"
  },
  "pickup": {
    "deadline": "int (days)",
    "location": "string",
    "hours": "string (optional)",
    "contact": "string (optional)"
  },
  "categories": ["array of strings"],
  "createdAt": "ISO 8601",
  "updatedAt": "ISO 8601"
}
```

### ⚠️ Braki w Zgodności z dane.gov.pl:

1. **Brak Metadanych DCAT**
2. **Brak @context URI**
3. **Brak linków (self, canonical)**
4. **Brak informacji licencji**
5. **Brak pól dostępności (accessibility)**
6. **Brak pól dla RODO/GDPR**

---

## 3. REKOMENDOWANY FORMAT (dane.gov.pl compatible)

### Opcja A: Minimalny Format (Rekomendowany)

```json
{
  "@context": "https://www.w3.org/ns/dcat",
  "@type": "dcat:Dataset",
  "id": "uuid",
  "dct:identifier": "uuid",
  "dct:title": "Portfel skórzany brązowy",
  "dct:description": "Brązowy portfel ze skóry naturalnej",
  "dcat:keyword": ["dokumenty", "portfele", "dowód osobisty"],

  "municipality": {
    "name": "Warszawa",
    "type": "miasto",
    "contactEmail": "kontakt@um.warszawa.pl",
    "organizationId": "http://www.wikidata.org/entity/Q270" // Wikidata URI
  },

  "item": {
    "name": "Portfel skórzany brązowy",
    "category": "dokumenty",
    "date": "2025-12-01",
    "location": "Park Łazienkowski, ławka przy fontannie",
    "status": "available",
    "description": "Brązowy portfel ze skóry naturalnej, zawiera dowód osobisty"
  },

  "pickup": {
    "deadline": 30,
    "location": "Urząd Miasta Warszawa, ul. Senatorska 30",
    "hours": "Poniedziałek-Piątek 9:00-17:00",
    "contact": "+48 22 443 44 44",
    "method": "personal" // new
  },

  "dct:issued": "2025-12-01T10:00:00Z",
  "dct:modified": "2025-12-01T10:00:00Z",
  "dcat:landingPage": "https://zguba.gov/item/uuid",
  "dct:license": "http://creativecommons.org/licenses/by/4.0/",
  "dcat:distribution": {
    "dcat:accessURL": "https://api.zguba.gov/api/found-items/uuid"
  },
  "dcat:contactPoint": {
    "vcard:fn": "Urząd Miasta Warszawa",
    "vcard:hasEmail": "kontakt@um.warszawa.pl"
  }
}
```

### Opcja B: OData Compliant (dla dane.gov.pl API)

```json
{
  "odata.type": "Zguba.Models.FoundItem",
  "id": "uuid",
  "municipalityName": "Warszawa",
  "municipalityType": "miasto",
  "municipalityEmail": "kontakt@um.warszawa.pl",
  "itemName": "Portfel skórzany brązowy",
  "itemCategory": "dokumenty",
  "itemDate": "2025-12-01",
  "itemLocation": "Park Łazienkowski",
  "itemStatus": "available",
  "pickupDeadline": 30,
  "pickupLocation": "Urząd Miasta Warszawa",
  "createdAt": "2025-12-01T10:00:00Z",
  "updatedAt": "2025-12-01T10:00:00Z",
  "__metadata": {
    "type": "Zguba.Models.FoundItem",
    "uri": "https://api.zguba.gov/api/found-items('uuid')"
  }
}
```

---

## 4. REKOMENDACJE IMPLEMENTACJI

### 🔧 Co zmienić w Backend?

#### 1. Dodaj Metadane DCAT do Schematu

```python
# schemas/found_item.py - ZMIANA

class FoundItemResponse(BaseModel):
    # Existing fields
    id: str
    municipality: MunicipalityInfo
    item: ItemInfo
    pickup: PickupInfo
    categories: List[str]
    createdAt: Optional[str] = None
    updatedAt: Optional[str] = None

    # NOWE POLA DLA dane.gov.pl
    context: str = "https://www.w3.org/ns/dcat"  # @context
    type: str = "dcat:Dataset"  # @type
    license: str = "http://creativecommons.org/licenses/by/4.0/"  # Default license
    landing_page: Optional[str] = None  # Landing page URL

    class Config:
        from_attributes = True
        json_schema_extra = {
            "example": {
                "id": "550e8400-e29b-41d4-a716-446655440000",
                "context": "https://www.w3.org/ns/dcat",
                "type": "dcat:Dataset",
                # ... reszta
            }
        }
```

#### 2. Dodaj Endpoint OData

```python
# routers/found_items.py - NOWY ENDPOINT

@router.get("/odata", tags=["odata"])
async def get_found_items_odata(
    skip: int = Query(0, ge=0),
    limit: int = Query(50, ge=1, le=100),
    filter: Optional[str] = Query(None),
    db: AsyncSession = Depends(get_db)
):
    """
    OData endpoint dla dane.gov.pl
    Example: GET /odata?$filter=itemStatus eq 'available'&$top=10&$skip=0
    """
    # Implementation with OData filtering
    pass
```

#### 3. Dodaj CORS dla dane.gov.pl

```python
# main.py - ZMIANA CORS

app.add_middleware(
    CORSMiddleware,
    allow_origins=[
        "http://localhost:4200",
        "http://localhost:8000",
        "https://dane.gov.pl",
        "https://api.dane.gov.pl",
        "https://data.europa.eu",  # CKAN compatibility
    ],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)
```

#### 4. Dodaj Endpoint Metadanych

```python
# routers/metadata.py - NOWY PLIK

@router.get("/metadata", tags=["metadata"])
async def get_dataset_metadata():
    """
    Zwraca metadane datasetu w formacie DCAT-AP
    """
    return {
        "@context": "https://www.w3.org/ns/dcat",
        "@type": "dcat:Catalog",
        "dct:title": "Katalog Rzeczy Znalezionych - Zguba.gov",
        "dct:description": "System zarabiania znalezionych przedmiotów w Polsce",
        "dct:issued": "2025-12-01T00:00:00Z",
        "dct:modified": "2025-12-01T00:00:00Z",
        "dcat:dataset": [
            {
                "@type": "dcat:Dataset",
                "dct:title": "Znalezione Rzeczy",
                "dct:description": "Bieżąca lista znalezionych przedmiotów",
                "dcat:distribution": [
                    {
                        "@type": "dcat:Distribution",
                        "dct:title": "JSON API",
                        "dcat:accessURL": "https://api.zguba.gov/api/found-items",
                        "dct:format": "JSON",
                        "dcat:mediaType": "application/json"
                    }
                ]
            }
        ]
    }
```

---

## 5. INTEGRACJA Z dane.gov.pl

### Kroki Rejestracji:

1. Przejdź do: https://dane.gov.pl/katalog
2. Zaloguj się jako administrator
3. "Dodaj nowy zbiór danych"
4. Podaj URL naszego metadata endpoint: `https://api.zguba.gov/metadata`
5. dane.gov.pl zindeksuje dane

### Wymagane URI dla dane.gov.pl:

```
Katalog: https://dane.gov.pl/dataset/rzeczy-znalezione
API Endpoint: https://api.zguba.gov/api/found-items
Metadata: https://api.zguba.gov/metadata
OData: https://api.zguba.gov/odata (opcjonalnie)
```

---

## 6. TESTOWANIE ZGODNOŚCI

### Sprawdź Validację JSON-LD

```bash
# Test 1: Validacja @context
curl -H "Accept: application/ld+json" \
  https://api.zguba.gov/api/found-items/1 | jq '.["@context"]'

# Test 2: DCAT Compliance
curl https://api.zguba.gov/metadata | jq '.["@type"]'

# Test 3: OData Compliance
curl "https://api.zguba.gov/odata?$filter=status%20eq%20%27available%27"
```

### Validator DCAT-AP

- https://data.europa.eu/dcat-ap/validation
- https://www.w3.org/RDF/Distiller/ (RDF validation)

---

## 7. PODSUMOWANIE DZIAŁAŃ

### Priorytet 1 (Wymagane):

- ✅ Obecny format JSON jest OK dla podstawowych operacji
- ⚠️ Dodaj `@context` i `@type` do odpowiedzi API
- ⚠️ Dodaj pole `license` (CC-BY 4.0)

### Priorytet 2 (Rekomendowane):

- ⚠️ Zaimplementuj Metadata Endpoint
- ⚠️ Dodaj landing_page URL dla każdego przedmiotu
- ⚠️ Rozszerz CORS dla dane.gov.pl

### Priorytet 3 (Opcjonalnie):

- ⚠️ OData Endpoint
- ⚠️ RDF/Turtle Export
- ⚠️ CKAN Integration

---

## 8. PRZYKŁAD PEŁNEJ ODPOWIEDZI (Rekomendowany Format)

```json
{
  "@context": "https://www.w3.org/ns/dcat",
  "@type": "dcat:Dataset",
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "dct:identifier": "550e8400-e29b-41d4-a716-446655440000",
  "dct:title": "Portfel skórzany brązowy",
  "dct:description": "Brązowy portfel ze skóry naturalnej zawierający dowód osobisty",
  "dct:issued": "2025-12-01T10:00:00Z",
  "dct:modified": "2025-12-01T10:00:00Z",
  "dct:license": "http://creativecommons.org/licenses/by/4.0/",
  "dcat:keyword": ["dokumenty", "portfele", "dowód osobisty"],
  "dcat:landingPage": "https://zguba.gov/item/550e8400-e29b-41d4-a716-446655440000",

  "dcat:contactPoint": {
    "@type": "vcard:Organization",
    "vcard:fn": "Urząd Miasta Warszawa",
    "vcard:hasEmail": "mailto:kontakt@um.warszawa.pl",
    "vcard:hasTelephone": "+48 22 443 44 44"
  },

  "dcat:distribution": [
    {
      "@type": "dcat:Distribution",
      "dcat:accessURL": "https://api.zguba.gov/api/found-items/550e8400-e29b-41d4-a716-446655440000",
      "dct:format": "JSON",
      "dcat:mediaType": "application/json"
    }
  ],

  "municipality": {
    "name": "Warszawa",
    "type": "miasto",
    "contactEmail": "kontakt@um.warszawa.pl"
  },

  "item": {
    "name": "Portfel skórzany brązowy",
    "category": "dokumenty",
    "date": "2025-12-01",
    "location": "Park Łazienkowski, ławka przy fontannie",
    "status": "available",
    "description": "Brązowy portfel ze skóry naturalnej, zawiera dowód osobisty"
  },

  "pickup": {
    "deadline": 30,
    "location": "Urząd Miasta Warszawa, ul. Senatorska 30, pokój 215",
    "hours": "Poniedziałek-Piątek 9:00-17:00",
    "contact": "+48 22 443 44 44",
    "method": "personal"
  },

  "categories": ["dokumenty", "portfele"],
  "createdAt": "2025-12-01T10:00:00Z",
  "updatedAt": "2025-12-01T10:00:00Z"
}
```

---

## Kontakt: dane.gov.pl Support

- Portal: https://dane.gov.pl
- API Docs: https://api.dane.gov.pl/doc
- Support: support@dane.gov.pl
