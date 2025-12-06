# Autouzupełnianie Jednostek Terytorialnych z API GUS

## Opis funkcjonalności

Dodano system autouzupełniania dla pola "Nazwa samorządu" w pierwszym kroku formularza. System pobiera **aktualne dane z oficjalnego API GUS (Głównego Urzędu Statystycznego)** poprzez Bank Danych Lokalnych (BDL).

## Źródło danych

### API GUS BDL
- **URL**: `https://bdl.stat.gov.pl/api/v1`
- **Format**: JSON
- **Dokumentacja**: https://api.stat.gov.pl/Home/BdlApi

### Dane pobierane automatycznie:
- ✅ **16 województw** (Level 2)
- ✅ **~382 powiaty** (Level 5, kind=1)
- ✅ **~73 miasta na prawach powiatu** (Level 5, kind=2)
- ✅ **~500 gmin** (Level 6, pierwsza partia dla wydajności)

## Jak to działa?

1. **Automatyczne ładowanie przy starcie** - Dane pobierają się w tle przy inicjalizacji aplikacji
2. **Wybierz typ samorządu** - Użytkownik wybiera typ (Województwo, Powiat, Gmina lub Miasto)
3. **Zacznij pisać nazwę** - Po wpisaniu minimum 2 znaków system wyświetla pasujące sugestie z API
4. **Wybierz z listy** - Kliknięcie na sugestię automatycznie wypełnia formularz

## Przykłady użycia

### Przykład 1: Wyszukiwanie miasta na prawach powiatu
1. Wybierz typ: "Miasto"
2. Wpisz: "Warsz"
3. System pokaże: "Warszawa (woj. Mazowieckie)"

### Przykład 2: Wyszukiwanie powiatu
1. Wybierz typ: "Powiat"
2. Wpisz: "krak"
3. System pokaże wszystkie powiaty zawierające "krak", np.:
   - "Powiat krakowski (woj. Małopolskie)"
   - "Powiat krapkowicki (woj. Opolskie)"

### Przykład 3: Wyszukiwanie województwa
1. Wybierz typ: "Województwo"
2. Wpisz: "maz"
3. System pokaże: "Województwo Mazowieckie"

## Struktura techniczna

### API Endpoints GUS BDL:

```
GET https://bdl.stat.gov.pl/api/v1/units?level=2&format=json&page-size=100
→ Województwa (16 jednostek)

GET https://bdl.stat.gov.pl/api/v1/units?level=5&format=json&page-size=100&page={X}
→ Powiaty i miasta na prawach powiatu (~455 jednostek)
  - kind="1" → Powiaty
  - kind="2" → Miasta na prawach powiatu

GET https://bdl.stat.gov.pl/api/v1/units?level=6&format=json&page-size=500&page=0
→ Gminy (pierwsza partia ~500 z ~4180 jednostek)
```

### Pliki dodane:
- `src/app/services/territorial-units.service.ts` - Serwis zarządzający danymi z API GUS

### Pliki zmodyfikowane:
- `src/app/app.component.ts` - Dodano logikę autouzupełniania (async)
- `src/app/app.component.html` - Dodano UI autouzupełniania
- `src/app/app.component.css` - Dodano style dla autouzupełniania
- `src/app/found-items/found-items.component.ts` - Dodano logikę autouzupełniania (async)
- `src/app/found-items/found-items.component.html` - Dodano UI autouzupełniania
- `src/app/found-items/found-items.component.css` - Dodano style dla autouzupełniania

## Interface TerritorialUnit

```typescript
interface TerritorialUnit {
  name: string;                          // Nazwa jednostki
  type: 'wojewodztwo' | 'powiat' | 'gmina' | 'miasto';
  parentName?: string;                   // Województwo nadrzędne
  fullName: string;                      // Pełna nazwa z kontekstem
  id?: string;                           // ID TERYT z API GUS
}
```

## Metody API serwisu

### `async search(query: string, type?: TerritorialUnit['type']): Promise<TerritorialUnit[]>`
Wyszukuje jednostki na podstawie zapytania (asynchronicznie).

**Parametry:**
- `query` - Fraza do wyszukania (minimum 2 znaki)
- `type` - Opcjonalny filtr typu jednostki

**Zwraca:** Promise z tablicą pasujących jednostek (max 20)

### `getByType(type: TerritorialUnit['type']): TerritorialUnit[]`
Pobiera wszystkie jednostki określonego typu.

**Parametry:**
- `type` - Typ jednostki ('wojewodztwo', 'powiat', 'gmina', 'miasto')

**Zwraca:** Tablicę wszystkich jednostek danego typu

## Obsługa błędów i backup

System posiada mechanizm fallback:
- Jeśli API GUS nie odpowiada, ładowane są dane backup
- Dane backup zawierają podstawowy zestaw jednostek
- Timeout ładowania: 10 sekund
- Console logs informują o statusie ładowania

## Wydajność

- **Lazy loading** - Dane ładują się w tle przy starcie aplikacji
- **Caching** - Dane są cache'owane w pamięci po pierwszym załadowaniu
- **Pagination** - API używa stronicowania aby nie przeciążać przeglądarki
- **Limit gmin** - Domyślnie ładowanych jest 500 gmin (można rozszerzyć)

## Możliwości rozbudowy

### Pełne ładowanie gmin
Można rozszerzyć `loadGminy()` aby pobierać wszystkie ~4180 gmin:

```typescript
private async loadGminy(): Promise<void> {
  let page = 0;
  let hasMore = true;
  const pageSize = 500;
  
  while (hasMore) {
    const url = `${this.API_BASE}/units?level=6&format=json&page-size=${pageSize}&page=${page}`;
    const response = await fetch(url);
    const data: GusApiResponse = await response.json();
    
    // przetwarzanie...
    
    page++;
    hasMore = data.links?.next !== undefined;
  }
}
```

### Cache w localStorage
Można dodać cache'owanie w localStorage aby uniknąć częstych zapytań do API:

```typescript
private loadFromCache(): boolean {
  const cached = localStorage.getItem('gus_units_cache');
  if (cached) {
    const data = JSON.parse(cached);
    if (Date.now() - data.timestamp < 24 * 60 * 60 * 1000) { // 24h
      this.units = data.units;
      return true;
    }
  }
  return false;
}
```

## Uwagi

- ✅ Dane są pobierane z **oficjalnego API GUS**
- ✅ System działa **asynchronicznie** (nie blokuje UI)
- ✅ Posiada **mechanizm fallback** na wypadek problemów z API
- ✅ **Wszystkie województwa** i **powiaty** są ładowane automatycznie
- ⚠️ Gminy ładowane są częściowo dla wydajności (można rozszerzyć)
- 🌐 Wymaga **połączenia z internetem** przy pierwszym użyciu

## Limity API GUS

API GUS BDL jest publiczne i bezpłatne, ale:
- Może mieć limity zapytań (nie są publicznie udokumentowane)
- W przypadku problemów system używa danych backup
- Zalecane jest cache'owanie wyników
