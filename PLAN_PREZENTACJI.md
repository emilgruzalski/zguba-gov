# Plan Prezentacji - System Zguba.gov
## Prototyp "jednego okna" do udostępniania danych o rzeczach znalezionych

---

## 📋 Struktura Prezentacji (min. 5 slajdów)

---

### **Slajd 1: Strona tytułowa i Wprowadzenie**

**Treść:**
- Tytuł: **Zguba.gov - System udostępniania danych o rzeczach znalezionych**
- Podtytuł: Prototyp "jednego okna" zgodny z dane.gov.pl
- Data: Grudzień 2025
- Autorzy/zespół

**Główne punkty:**
- Problem: brak zunifikowanego systemu zgłaszania rzeczy znalezionych przez samorządy
- Rozwiązanie: centralna platforma zgodna z standardami dane.gov.pl
- Cel: max. 5 kroków do udostępnienia danych

---

### **Slajd 2: Architektura systemu (dowód techniczny #1)**

**Treść:**
- Diagram architektury z komponentami:
  - **Frontend**: Angular 19 (responsywny UI, WCAG 2.1)
  - **Backend**: FastAPI (Python) - REST API + OData
  - **Baza danych**: SQLite/PostgreSQL
  - **Integracja**: endpoint OData zgodny z dane.gov.pl

**Zrzut ekranu:**
- Schemat folderów projektu
- Widok struktury kodu (backend/frontend)
- Screenshot z dokumentacji API (Swagger UI na /docs)

**Kluczowe punkty:**
- ✅ Format odczytywalny maszynowo (JSON, OData)
- ✅ REST API + OData dla kompatybilności z dane.gov.pl
- ✅ Schemat DCAT dla metadanych

---

### **Slajd 3: Proces zgłoszenia w 5 krokach (dowód UX #2)**

**Treść: "Ścieżka urzędnika - 5 kroków do publikacji danych"**

**KROK 1:** Wybór jednostki terytorialnej
- Zrzut ekranu: Lista gmin/powiatów z wyszukiwarką
- Autocomplete z pełną listą JST

**KROK 2:** Wprowadzenie podstawowych danych o rzeczy
- Formularz: nazwa, kategoria, data znalezienia, lokalizacja
- Walidacja pól

**KROK 3:** Dodanie szczegółów odbioru
- Termin odbioru (dni), miejsce, godziny, kontakt
- Autouzupełnianie danych urzędu

**KROK 4:** Przypisanie kategorii/tagów
- Wielokrotny wybór kategorii (dokumenty, elektronika, klucze, etc.)
- Sugerowane tagi

**KROK 5:** Publikacja
- Podsumowanie danych
- Przycisk "Opublikuj" → automatyczne udostępnienie przez API

**Zrzuty ekranu:**
- Screenshot każdego kroku formularza
- Widok mobilny + desktop (responsywność)

---

### **Slajd 4: Zgodność z dane.gov.pl (dowód techniczny #3)**

**Treść:**
- **Format danych**: JSON (odczytywalny maszynowo) ❌ PDF ❌ DOC
- **Endpoint OData**: `/odata` - zgodny z protokołem OData v2/v4
- **Metadane DCAT**: 
  - `dct:identifier`, `dct:title`, `dct:description`
  - `dct:issued`, `dct:modified`, `dct:license`
  - `dcat:keyword`, `dcat:landingPage`

**Zrzut ekranu:**
- Przykład odpowiedzi JSON z API (`/api/found-items`)
- Endpoint OData metadata (`/odata/$metadata`)
- Przykład zapytania OData z filtrami:
  ```
  /odata?$filter=item_status eq 'available'&$top=10
  ```

**Kluczowe punkty:**
- ✅ Pełna zgodność z standardem DCAT-AP
- ✅ Wsparcie dla OData query options ($filter, $top, $skip, $orderby)
- ✅ Licencja Creative Commons BY 4.0

---

### **Slajd 5: Funkcjonalności systemu (dowód funkcjonalny #4)**

**Treść:**
- **Dla urzędników:**
  - ✅ Dodawanie rzeczy znalezionych (formularz 5-krokowy)
  - ✅ Edycja i aktualizacja statusu (available → claimed → expired)
  - ✅ Zarządzanie kategoriami
  
- **Dla obywateli:**
  - 🔍 Wyszukiwanie po kategorii, gminie, dacie
  - 📊 Przeglądanie listy rzeczy znalezionych
  - 📍 Filtrowanie po lokalizacji
  - 📱 Dostęp mobilny

- **Dla dane.gov.pl:**
  - 🔌 API REST + OData
  - 📋 Eksport danych w JSON
  - 📈 Statystyki agregowane

**Zrzuty ekranu:**
- Lista rzeczy znalezionych (widok główny)
- Filtry i wyszukiwarka
- Panel statystyk (`/api/stats`)
- Widok mobilny (responsywność)

---

### **Slajd 6: Dostępność i responsywność (dowód WCAG/RWD #5)**

**Treść:**
- **WCAG 2.1 Compliance:**
  - ✅ Kontrast kolorów (minimum AA)
  - ✅ Nawigacja klawiaturą (focus indicators)
  - ✅ Opisy alternatywne (ARIA labels)
  - ✅ Semantyczny HTML
  - ✅ Testowane screen readerem

- **Responsywność (RWD):**
  - ✅ Desktop (1920px+)
  - ✅ Tablet (768px-1024px)
  - ✅ Mobile (320px-767px)
  - ✅ Elastyczne layouty (Flexbox/Grid)

**Zrzuty ekranu:**
- Widok na 3 urządzeniach jednocześnie (desktop/tablet/mobile)
- Test kontrastu kolorów
- Nawigacja klawiaturą (focus states)

---

### **Slajd 7: Demo i technologie**

**Treść:**
- **Link do działającego demo:**
  - Frontend: `http://localhost:4200`
  - Backend API: `http://localhost:8000`
  - Dokumentacja: `http://localhost:8000/docs`
  - OData endpoint: `http://localhost:8000/odata`

- **Wideo demo (2-3 min):** 
  - Pokazanie pełnej ścieżki użytkownika (5 kroków)
  - Demonstracja API (Postman/curl)
  - Responsywność na różnych urządzeniach

- **Stack technologiczny:**
  - Frontend: Angular 19, TypeScript, CSS3
  - Backend: Python 3.9+, FastAPI, SQLAlchemy
  - Baza: SQLite (dev) / PostgreSQL (prod)
  - API: REST + OData v2/v4

---

### **Slajd 8: Ograniczenia i dalszy rozwój**

**Treść:**
- **Ograniczenia prototypu:**
  - ⚠️ Brak autentykacji/autoryzacji (OAuth2 w kolejnej wersji)
  - ⚠️ Lokalna baza SQLite (migracja do PostgreSQL dla produkcji)
  - ⚠️ Brak integracji z ePUAP/profil zaufany
  - ⚠️ Przykładowe dane jednostek terytorialnych (JSON statyczny)
  - ⚠️ Brak automatycznego powiadamiania (email/SMS)

- **Planowany rozwój:**
  - 🔐 System logowania dla urzędników (OAuth2 + JWT)
  - 📧 Powiadomienia email/SMS o nowych rzeczach
  - 🔗 Bezpośrednia integracja z katalogiem dane.gov.pl
  - 🗺️ Mapa interaktywna z miejscami znalezienia
  - 📊 Dashboard analityczny dla JST
  - 🌐 Wielojęzyczność (PL/EN)

---

## 📦 Pliki do dostarczenia

### 1. **Prezentacja (format PDF/PPTX)**
   - Min. 8 slajdów zgodnie z powyższym planem
   - Zrzuty ekranu z każdego kroku użytkownika
   - Screenshoty API, kodu, struktury projektu

### 2. **Działający prototyp**
   - **Demo live:** 
     - Link do wdrożonego systemu (np. Vercel + Render/Railway)
     - LUB instrukcja uruchomienia lokalnego
   
   - **Wideo (2-3 min):**
     - Nagranie ekranu pokazujące:
       1. Proces dodawania rzeczy (5 kroków)
       2. Wyszukiwanie i filtrowanie
       3. Wywołanie API (Postman/curl)
       4. Responsywność (zmiana rozdzielczości)
       5. Dostępność (nawigacja klawiaturą)

### 3. **Wzorcowy zakres danych (format JSON)**
   
   **Plik:** `example_item_schema.json`
   
   ```json
   {
     "@context": "https://www.w3.org/ns/dcat",
     "@type": "dcat:Dataset",
     "dct:identifier": "uuid-v4",
     "dct:title": "Nazwa znalezionej rzeczy",
     "dct:description": "Szczegółowy opis rzeczy",
     "dct:issued": "2025-12-07T10:00:00Z",
     "dct:modified": "2025-12-07T10:00:00Z",
     "dct:license": "http://creativecommons.org/licenses/by/4.0/",
     "dcat:keyword": ["kategoria1", "kategoria2"],
     "dcat:landingPage": "https://zguba.gov/items/{id}",
     
     "municipality": {
       "name": "Warszawa",
       "type": "miasto na prawach powiatu",
       "teryt": "1465011",
       "contactEmail": "kontakt@um.warszawa.pl",
       "contactPhone": "+48 22 443 44 44"
     },
     
     "item": {
       "name": "Portfel skórzany",
       "category": "dokumenty",
       "foundDate": "2025-12-01",
       "foundLocation": "Park Łazienkowski, alejka główna",
       "status": "available",
       "description": "Brązowy portfel ze skóry zawierający dokumenty"
     },
     
     "pickup": {
       "deadlineDays": 30,
       "location": "Urząd Miasta Warszawa, ul. Senatorska 30, 00-082 Warszawa",
       "hours": "Poniedziałek-Piątek 9:00-17:00",
       "contact": "+48 22 443 44 44",
       "method": "personal",
       "notes": "Prosimy o wcześniejszy kontakt telefoniczny"
     },
     
     "metadata": {
       "createdAt": "2025-12-07T10:00:00Z",
       "updatedAt": "2025-12-07T10:00:00Z",
       "createdBy": "urząd.warszawa",
       "dataVersion": "1.0"
     }
   }
   ```

### 4. **Kod źródłowy**
   - **Repository:** GitHub/GitLab (publiczny)
   - **Struktura:**
     ```
     zguba-gov/
     ├── README.md                 # Instrukcja instalacji
     ├── PLAN_PREZENTACJI.md       # Ten dokument
     ├── example_item_schema.json  # Wzorcowy format danych
     ├── backend/
     │   ├── requirements.txt
     │   ├── main.py
     │   ├── models/
     │   ├── routers/
     │   └── schemas/
     └── frontend/
         ├── package.json
         └── src/
     ```
   - **Licencja:** MIT / Apache 2.0
   - **Dokumentacja:**
     - README z instrukcją uruchomienia
     - API_DOCS.md z opisem endpointów
     - Komentarze w kodzie

---

## 🎯 Checklist przed prezentacją

- [ ] **Prezentacja (min. 5 slajdów):**
  - [ ] Zrzuty ekranu każdego kroku (5 kroków użytkownika)
  - [ ] Screenshot struktury projektu
  - [ ] Screenshot API documentation (Swagger)
  - [ ] Screenshot OData endpoint
  - [ ] Widoki mobile/desktop (responsywność)
  - [ ] Dowód WCAG (kontrast, focus)

- [ ] **Prototyp:**
  - [ ] System działa lokalnie (backend + frontend)
  - [ ] API zwraca dane w JSON
  - [ ] Endpoint OData działa
  - [ ] Formularz 5-krokowy działa
  - [ ] Wyszukiwanie/filtrowanie działa
  - [ ] Responsywność (mobile/tablet/desktop)

- [ ] **Wideo (2-3 min):**
  - [ ] Nagranie procesu dodawania rzeczy
  - [ ] Demonstracja API
  - [ ] Test responsywności
  - [ ] Test dostępności (klawiatura)

- [ ] **Wzorcowy zakres danych:**
  - [ ] Plik `example_item_schema.json`
  - [ ] Zgodność z DCAT
  - [ ] Wszystkie wymagane pola
  - [ ] Przykładowe wartości

- [ ] **Kod źródłowy:**
  - [ ] Repository na GitHubie
  - [ ] README z instrukcją
  - [ ] Kod skomentowany
  - [ ] requirements.txt / package.json
  - [ ] Skrypt inicjalizujący bazę

---

## 📊 Kluczowe metryki do pokazania

### Statystyki systemu (endpoint `/api/stats`):
```json
{
  "totalItems": 150,
  "availableItems": 120,
  "claimedItems": 25,
  "expiredItems": 5,
  "topCategories": [
    {"category": "dokumenty", "count": 45},
    {"category": "klucze", "count": 30},
    {"category": "elektronika", "count": 25}
  ],
  "topMunicipalities": [
    {"name": "Warszawa", "count": 50},
    {"name": "Kraków", "count": 30}
  ]
}
```

### Przykładowe zapytania OData:
```
# Wszystkie dostępne rzeczy
/odata?$filter=item_status eq 'available'

# Dokumenty w Warszawie
/odata?$filter=item_category eq 'dokumenty' and municipality_name eq 'Warszawa'

# Ostatnie 10 rzeczy
/odata?$orderby=created_at desc&$top=10

# Z licznikiem
/odata?$count=true&$top=20
```

---

## 🎬 Scenariusz wideo (2-3 min)

**00:00-00:15** - Intro
- Ekran startowy systemu
- Krótkie wprowadzenie głosowe

**00:15-01:30** - Proces 5-krokowy (główna funkcjonalność)
- KROK 1: Wybór Warszawa z listy
- KROK 2: Wypełnienie danych: "Portfel skórzany", kategoria "dokumenty"
- KROK 3: Miejsce odbioru, termin 30 dni
- KROK 4: Dodanie tagów
- KROK 5: Klik "Opublikuj" → potwierdzenie

**01:30-02:00** - Wyszukiwanie i API
- Wyszukanie dodanej rzeczy
- Otwarcie Postman/curl
- Wywołanie `/api/found-items` → JSON response
- Wywołanie `/odata?$filter=...` → OData response

**02:00-02:30** - Responsywność i dostępność
- Zmiana rozmiaru okna (desktop → mobile)
- Nawigacja TAB (focus indicators)
- Screen reader (opcjonalnie)

**02:30-03:00** - Podsumowanie
- Wyświetlenie statystyk
- Link do dokumentacji
- Ekran końcowy

---

## 🔗 Przydatne linki

- **Dane.gov.pl API:** https://dane.gov.pl/api/
- **DCAT-AP standard:** https://www.w3.org/TR/vocab-dcat-2/
- **OData v4:** https://www.odata.org/documentation/
- **WCAG 2.1:** https://www.w3.org/WAI/WCAG21/quickref/
- **FastAPI docs:** https://fastapi.tiangolo.com/
- **Angular docs:** https://angular.dev/

---

## 📝 Notatki końcowe

### Mocne strony projektu:
✅ Prosty 5-krokowy proces (UX)  
✅ Pełna zgodność z dane.gov.pl (OData + DCAT)  
✅ Format JSON (machine-readable)  
✅ REST API + dokumentacja Swagger  
✅ Responsywność (mobile-first)  
✅ Podstawowa dostępność WCAG  
✅ Open source (kod dostępny)  

### Obszary do podkreślenia:
🎯 **Prosty UI** - urzędnik nie potrzebuje szkolenia  
🎯 **Automatyzacja** - dane automatycznie dostępne przez API  
🎯 **Standaryzacja** - jeden schemat dla całej Polski  
🎯 **Interoperacyjność** - łatwa integracja z innymi systemami  

---

**Data utworzenia planu:** 7 grudnia 2025  
**Wersja:** 1.0  
**Status:** Gotowy do realizacji
