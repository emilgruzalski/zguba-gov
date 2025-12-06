# Lokalna baza danych jednostek terytorialnych

## Zmiana z API GUS na lokalny plik

Aplikacja została zmieniona z pobierania danych z API GUS na **korzystanie z lokalnego pliku Excel** zawierającego pełną bazę jednostek samorządu terytorialnego (JST).

## Plik źródłowy

**Lokalizacja:** `db/baza_jst_aktualizacja_26062025.xlsx`

Plik zawiera:
- **2827 rekordów** jednostek samorządu terytorialnego
- **Rzeczywiste adresy email** jednostek
- **Pełne dane kontaktowe** (nazwa urzędu, adres, telefon, www)
- **Aktualny na dzień:** 26 czerwca 2025

## Struktura danych w pliku Excel

| Kolumna | Opis | Przykład |
|---------|------|----------|
| `Kod_TERYT` | Unikalny kod jednostki | `0201000` |
| `nazwa_samorządu` | Nazwa jednostki | `Powiat Bolesławiecki` |
| `Województwo` | Nazwa województwa | `dolnośląskie` |
| `Powiat` | Nazwa powiatu | `bolesławiecki` |
| `typ_JST` | Typ jednostki | `W`, `P`, `MNP`, `GW`, `GM`, `GMW` |
| `nazwa_urzędu_JST` | Nazwa urzędu | `Starostwo Powiatowe w Bolesławcu` |
| `ogólny adres poczty elektronicznej...` | **Email kontaktowy** | `biuro@powiatboleslawiecki.pl` |
| `adres www jednostki` | Strona www | `www.powiatboleslawiecki.pl` |

## Typy JST

| Kod | Pełna nazwa | Liczba jednostek |
|-----|-------------|------------------|
| `W` | Województwo | 16 |
| `P` | Powiat | 314 |
| `MNP` | Miasto Na Prawach powiatu | 66 |
| `GW` | Gmina Wiejska | 1457 |
| `GM` | Gmina Miejska | 238 |
| `GMW` | Gmina Miejsko-Wiejska | 718 |
| `dzielnica` | Dzielnica (pomijana) | 18 |

## Konwersja do JSON

Aplikacja używa pliku JSON wygenerowanego z Excel:

**Lokalizacja:** `src/assets/territorial-units.json`

### Uruchomienie konwersji

```bash
npm run convert-excel
```

Lub bezpośrednio:

```bash
node scripts/convert-excel-to-json.js
```

### Proces konwersji

1. **Odczyt pliku Excel** za pomocą biblioteki `xlsx`
2. **Mapowanie typów JST:**
   - `W` → `wojewodztwo`
   - `P` → `powiat`
   - `MNP` → `miasto`
   - `GW`, `GM`, `GMW` → `gmina`
3. **Pomijanie dzielnic** (typ `dzielnica`)
4. **Zachowanie rzeczywistych emaili** z kolumny pliku Excel
5. **Zapis do JSON** w katalogu `src/assets/`

## Dane w aplikacji

Po konwersji aplikacja zawiera:

```
✅ Załadowano łącznie 2809 jednostek terytorialnych
   - Województwa: 16
   - Powiaty: 314
   - Miasta: 66
   - Gminy: 2413
```

## Przykładowe dane

### Województwo
```json
{
  "id": "0200000",
  "name": "Województwo dolnośląskie",
  "type": "wojewodztwo",
  "email": "umwd@dolnyslask.pl",
  "officeName": "Urząd Marszałkowski Województwa Dolnośląskiego"
}
```

### Miasto na prawach powiatu
```json
{
  "id": "0461000",
  "name": "Bydgoszcz",
  "type": "miasto",
  "email": "urzad@um.bydgoszcz.pl; bczk@um.bydgoszcz.pl",
  "officeName": "Urząd Miasta Bydgoszczy"
}
```

### Powiat
```json
{
  "id": "0201000",
  "name": "Powiat Bolesławiecki",
  "type": "powiat",
  "email": "biuro@powiatboleslawiecki.pl",
  "officeName": "Starostwo Powiatowe w Bolesławcu",
  "voivodeship": "dolnośląskie"
}
```

### Gmina
```json
{
  "id": "0201022",
  "name": "Bolesławiec",
  "type": "gmina",
  "email": "urzadgminy@gminaboleslawiec.pl",
  "officeName": "Urząd Gminy Bolesławiec",
  "voivodeship": "dolnośląskie",
  "county": "bolesławiecki"
}
```

## Auto-uzupełnianie emaila

Aplikacja **priorytetowo używa rzeczywistych emaili z bazy danych**.

Jeśli email nie jest dostępny, generuje go algorytmicznie:
- Województwo → `kontakt@{slug}.uw.gov.pl`
- Powiat → `starostwo@{slug}.pl`
- Miasto → `urzad@um.{slug}.pl`
- Gmina → `ug@{slug}.pl`

### Przykłady rzeczywistych emaili

| Jednostka | Rzeczywisty email z bazy | Wygenerowany (fallback) |
|-----------|--------------------------|-------------------------|
| Bydgoszcz | `urzad@um.bydgoszcz.pl` | `urzad@um.bydgoszcz.pl` |
| Warszawa | `um@um.warszawa.pl` | `urzad@um.warszawa.pl` |
| Województwo mazowieckie | `kancelaria@mazovia.pl` | `kontakt@mazowieckie.uw.gov.pl` |
| Powiat krakowski | `starostwo@powiat-krakowski.pl` | `starostwo@krakowski.pl` |

✅ **Zaleta:** Rzeczywiste emaile są dokładniejsze i aktualne!

## Aktualizacja danych

Gdy pojawi się nowy plik Excel z GUS:

1. **Umieść plik** w katalogu `db/`
2. **Zaktualizuj ścieżkę** w `scripts/convert-excel-to-json.js`:
   ```javascript
   const excelPath = path.join(__dirname, '../db/NOWA_NAZWA_PLIKU.xlsx');
   ```
3. **Uruchom konwersję:**
   ```bash
   npm run convert-excel
   ```
4. **Sprawdź wynik** w konsoli (liczba jednostek, przykłady)
5. **Zrestartuj aplikację** Angular

## Zalety rozwiązania

✅ **Szybsze działanie** - brak wywołań API  
✅ **Rzeczywiste emaile** - z oficjalnej bazy GUS  
✅ **Offline** - działa bez internetu  
✅ **Pełne dane** - wszystkie 2809 jednostek JST  
✅ **Łatwa aktualizacja** - wymiana pliku Excel  
✅ **Dodatkowe dane** - nazwy urzędów, adresy www, itp.

## Technologie

- **xlsx** - biblioteka do odczytu plików Excel
- **TypeScript** - `resolveJsonModule: true` w `tsconfig.json`
- **Angular** - import statyczny pliku JSON
