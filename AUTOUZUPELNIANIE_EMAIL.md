# Autouzupełnianie Emaila Kontaktowego

## Funkcjonalność

Po wybraniu nazwy urzędu z autouzupełniania, pole **Email kontaktowy** automatycznie wypełnia się sugerowanym adresem email.

## Jak to działa?

1. Użytkownik wybiera urząd z listy autouzupełniania
2. System automatycznie generuje email na podstawie:
   - Nazwy jednostki
   - Typu jednostki (województwo, powiat, gmina, miasto)
3. Email jest edytowalny - użytkownik może go zmienić

## Schemat generowania emaili

### Województwa
- **Format**: `kontakt@{nazwa}.uw.gov.pl`
- **Przykład**: "Województwo Mazowieckie" → `kontakt@mazowieckie.uw.gov.pl`

### Powiaty
- **Format**: `starostwo@{nazwa}.pl`
- **Przykład**: "Powiat krakowski" → `starostwo@krakowski.pl`

### Miasta na prawach powiatu
- **Format**: `um@{nazwa}.pl`
- **Przykład**: "Kraków" → `um@krakow.pl`
- **Przykład**: "Warszawa" → `um@warszawa.pl`

### Gminy
- **Format**: `ug@{nazwa}.pl`
- **Przykład**: "Gmina Wieliczka" → `ug@wieliczka.pl`

## Normalizacja nazw

System automatycznie:
- Usuwa przedrostki ("Województwo", "Powiat", "Gmina", "Miasto")
- Zamienia polskie znaki na ASCII (ą→a, ć→c, ę→e, ł→l, ń→n, ó→o, ś→s, ź→z, ż→z)
- Zamienia spacje na myślniki
- Usuwa znaki specjalne
- Konwertuje na małe litery

## Przykłady

| Nazwa jednostki | Typ | Wygenerowany email |
|----------------|-----|-------------------|
| Kraków | Miasto | um@krakow.pl |
| Warszawa | Miasto | um@warszawa.pl |
| Województwo Mazowieckie | Województwo | kontakt@mazowieckie.uw.gov.pl |
| Województwo Małopolskie | Województwo | kontakt@malopolskie.uw.gov.pl |
| Powiat krakowski | Powiat | starostwo@krakowski.pl |
| Powiat wrocławski | Powiat | starostwo@wroclawski.pl |
| Gmina Wieliczka | Gmina | ug@wieliczka.pl |
| Gmina Konstancin-Jeziorna | Gmina | ug@konstancin-jeziorna.pl |

## UI/UX

### Wskazówka w interfejsie
Pod polem email wyświetla się podpowiedź:
```
✨ Email uzupełni się automatycznie po wybraniu urzędu (możesz edytować)
```

### Przepływ użytkownika
1. Wybierz typ samorządu
2. Zacznij pisać nazwę
3. Wybierz z listy autouzupełniania
4. ✅ Email automatycznie się wypełni
5. (Opcjonalnie) Edytuj email jeśli potrzeba

## Implementacja techniczna

### Nowa metoda w TerritorialUnitsService
```typescript
generateContactEmail(unit: TerritorialUnit): string
```

### Aktualizacja w komponentach
```typescript
selectUnit(unit: TerritorialUnit): void {
  const suggestedEmail = this.territorialUnitsService.generateContactEmail(unit);
  
  this.form.patchValue({
    municipalityName: unit.name,
    municipalityType: unit.type,
    contactEmail: suggestedEmail  // ✨ Nowe!
  });
}
```

## Pliki zmienione

- ✅ `src/app/services/territorial-units.service.ts`
  - Dodano metodę `generateContactEmail()`
  
- ✅ `src/app/app.component.ts`
  - Zaktualizowano `selectUnit()` aby wypełniać email
  
- ✅ `src/app/found-items/found-items.component.ts`
  - Zaktualizowano `selectUnit()` aby wypełniać email
  
- ✅ `src/app/app.component.html`
  - Dodano podpowiedź pod polem email
  
- ✅ `src/app/found-items/found-items.component.html`
  - Dodano podpowiedź pod polem email

## Uwagi

- ⚠️ Wygenerowane emaile to **sugestie** - mogą nie być prawdziwymi adresami
- ✅ Użytkownik **może i powinien** edytować email jeśli jest inny
- 💡 W przyszłości można rozważyć integrację z rzeczywistą bazą kontaktów urzędów
- 📝 Format emaili jest zgodny z konwencją polskich urzędów administracji publicznej:
  - `uw.gov.pl` - urzędy wojewódzkie
  - `um@miasto.pl` - urzędy miejskie
  - `ug@gmina.pl` - urzędy gminne
  - `starostwo@powiat.pl` - starostwa powiatowe
