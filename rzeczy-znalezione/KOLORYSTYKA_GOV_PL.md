# Aktualizacja kolorystyki - GOV.PL / dane.gov.pl

## Zmiana stylistyki

Aplikacja została dostosowana do oficjalnej kolorystyki portalu **dane.gov.pl** i systemu **GOV.PL**.

## Nowa paleta kolorów

### Kolory główne
```css
--gov-blue: #0052B4         /* Główny niebieski GOV.PL */
--gov-blue-dark: #003d87    /* Ciemniejszy odcień dla hover */
--gov-blue-light: #4a90e2   /* Jaśniejszy odcień dla akcentów */
--gov-red: #DC143C          /* Kolor akcent/alert */
```

### Kolory neutralne
```css
--gov-gray-bg: #F5F5F5      /* Tło strony */
--gov-gray-light: #E8E8E8   /* Lekkie obramowania */
--gov-gray-dark: #666666    /* Tekst pomocniczy */
--gov-text: #333333         /* Tekst główny */
--gov-white: #FFFFFF        /* Białe tło */
```

### Cienie
```css
--shadow-sm: 0 2px 4px rgba(0, 82, 180, 0.08)
--shadow-md: 0 4px 12px rgba(0, 82, 180, 0.12)
--shadow-lg: 0 8px 24px rgba(0, 82, 180, 0.16)
```

## Zmienione elementy

### 1. **Header / Nagłówek**
- **Było:** Gradient niebieski #0f3a7d → #1e5a96 z okrągłymi kształtami
- **Jest:** Jednolity kolor GOV.PL blue #0052B4 z ciemniejszym borderem

### 2. **Progress Bar**
- **Było:** Jasny gradient z zaokrąglonymi rogami
- **Jest:** Prosty pasek GOV.PL blue bez zaokrągleń

### 3. **Kroki (Steps)**
- **Było:** Jasne szare (#e2e8f0) z gradientem dla aktywnego (#0f3a7d → #2563eb)
- **Jest:** GOV.PL gray (#E8E8E8) i GOV.PL blue (#0052B4) dla aktywnego

### 4. **Przyciski**
- **Było:** 
  - Primary: Gradient #0f3a7d → #2563eb z dużym cieniem
  - Secondary: Szare tło z obramowaniem
- **Jest:**
  - Primary: Jednolity #0052B4 (GOV.PL blue)
  - Secondary: Przezroczyste z obramowaniem #0052B4

### 5. **Formularze**
- **Było:** 
  - Zaokrąglone rogi (10px)
  - Jasne tło #f8fafc
  - Focus: #2563eb
- **Jest:**
  - Kwadratowe rogi (4px)
  - Białe tło
  - Focus: #0052B4 (GOV.PL blue)
  - Obramowanie: #E8E8E8

### 6. **Autocomplete dropdown**
- **Było:** Zaokrąglony z niebieskimi tagami #0066cc
- **Jest:** Kwadratowy z tagami w kolorze GOV.PL blue #0052B4

### 7. **Summary / Podsumowanie**
- **Było:** Gradient tła z zaokrągloną ramką
- **Jest:** Jednolite szare tło #F9FAFB z prostą lewą ramką #0052B4

## Charakterystyka GOV.PL Design System

✅ **Minimalizm** - brak gradientów, proste formy  
✅ **Czytelność** - wysokie kontrasty, duże fonty  
✅ **Dostępność** - zgodność z WCAG  
✅ **Konsystencja** - jeden zestaw kolorów  
✅ **Profesjonalizm** - oficjalny wygląd rządowy  

## Porównanie

| Element | Przed | Po |
|---------|-------|-----|
| Główny kolor | #0f3a7d → #2563eb | #0052B4 |
| Zaokrąglenia | 10-16px | 0-4px |
| Cienie | Duże, kolorowe | Subtelne, niebieskie |
| Gradienty | Wszędzie | Brak |
| Styl | Nowoczesny, miękki | Oficjalny, rządowy |

## Pliki zmienione

1. ✅ `src/styles.css` - zmienne globalne CSS
2. ✅ `src/app/app.component.css` - główny komponent
3. ✅ `src/app/found-items/found-items.component.css` - komponent found-items

## Jak używać zmiennych CSS

W stylach komponentów możesz używać zmiennych GOV.PL:

```css
.my-element {
  background: var(--gov-blue);
  color: var(--gov-white);
  border: 2px solid var(--gov-gray-light);
  box-shadow: var(--shadow-md);
}
```

## Zgodność

Kolorystyka jest zgodna z:
- ✅ Portal **dane.gov.pl**
- ✅ System **GOV.PL**
- ✅ Wytyczne dostępności **WCAG 2.1**
- ✅ Standardy rządowe RP
