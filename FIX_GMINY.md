# Fix: Autouzupełnianie dla gmin

## Problem
Autouzupełnianie nie pokazywało żadnych gmin przy wyszukiwaniu.

## Przyczyna
1. **Błędna logika filtrowania** - Kod pomijał WSZYSTKIE gminy z `kind=4` i `kind=5`, ale nie uwzględniał gmin z `kind=1,2,3`
2. **Zbyt duży pageSize** - API GUS ma limit 100 rekordów na stronę, a kod próbował pobrać 500

## Rozwiązanie

### 1. Poprawiono logikę filtrowania gmin
W API GUS gminy mają następujące typy:
- `kind=1` - gmina miejska (np. "Bochnia")
- `kind=2` - gmina wiejska (np. "Drwinia")
- `kind=3` - gmina miejsko-wiejska (np. "Nowy Wiśnicz")
- `kind=4` - część miejska gminy miejsko-wiejskiej (np. "Nowy Wiśnicz - miasto")
- `kind=5` - część wiejska gminy miejsko-wiejskiej (np. "Nowy Wiśnicz - obszar wiejski")

**Poprawka**: Teraz pomijamy TYLKO `kind=4` i `kind=5` (części większych gmin), a ładujemy `kind=1,2,3` (właściwe gminy).

### 2. Dostosowano do limitów API
- Zmieniono `pageSize` z 500 na 100 (maksymalny limit API)
- Zwiększono liczbę stron z 3 do 15
- Wynik: ~960 gmin załadowanych do autouzupełniania

### 3. Dodano szczegółowe logi
```typescript
console.log(`✅ Załadowano ${this.units.filter(u => u.type === 'gmina').length} gmin`);
```

## Rezultat
Po zmianach, w konsoli przeglądarki powinny pojawić się logi:
```
🔄 Rozpoczynam ładowanie danych z API GUS...
✅ Załadowano ~960 gmin
✅ Załadowano łącznie ~1500 jednostek terytorialnych z API GUS
   - Województwa: 16
   - Powiaty: ~309
   - Miasta: ~73
   - Gminy: ~960
```

## Testowanie
1. Odśwież stronę w przeglądarce (Cmd+R)
2. Wybierz typ: "Gmina"
3. Wpisz np. "Leszno" lub "Wielicz"
4. Powinny pojawić się gminy z całej Polski

## Pliki zmienione
- `src/app/services/territorial-units.service.ts`
  - Poprawiono metodę `loadGminy()`
  - Dostosowano do limitów API (pageSize=100)
  - Dodano szczegółowe logi
  - Zwiększono liczbę ładowanych stron (15 zamiast 3)
