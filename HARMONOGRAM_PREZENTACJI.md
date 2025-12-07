# Harmonogram przygotowania prezentacji

## 🗓️ Plan działania - 7 dni do prezentacji

---

## Dzień 1-2: Przygotowanie materiałów wizualnych (SLAJDY)

### Zadania:
- [ ] Zrobić zrzuty ekranu wszystkich 5 kroków formularza
- [ ] Screenshot dokumentacji API (Swagger)
- [ ] Screenshot struktury projektu
- [ ] Screenshot odpowiedzi JSON z API
- [ ] Screenshot OData endpoint
- [ ] Zrzuty ekranu widoków mobile/desktop/tablet
- [ ] Test kontrastu kolorów (WAVE/axe DevTools)
- [ ] Screenshot nawigacji klawiaturą (focus states)

### Narzędzia:
- macOS: Cmd+Shift+4
- Chrome DevTools: Cmd+Shift+P → "Screenshot"
- WAVE extension: https://wave.webaim.org/extension/
- axe DevTools: https://www.deque.com/axe/devtools/

### Rezultat:
- Folder `/screenshots` z 15-20 zrzutami ekranu
- Nazwy plików: `01_krok1_wybor_jednostki.png`, `02_krok2_dane_rzeczy.png`, etc.

---

## Dzień 3: Stworzenie prezentacji (PowerPoint/Google Slides/Keynote)

### Zadania:
- [ ] Slajd 1: Strona tytułowa
- [ ] Slajd 2: Architektura (diagram + screenshoty)
- [ ] Slajd 3: Proces 5-krokowy (każdy krok z UI)
- [ ] Slajd 4: Zgodność z dane.gov.pl (JSON, OData, DCAT)
- [ ] Slajd 5: Funkcjonalności (lista + screenshoty)
- [ ] Slajd 6: WCAG i responsywność (3 urządzenia)
- [ ] Slajd 7: Demo i technologie (linki, stack)
- [ ] Slajd 8: Ograniczenia i rozwój (lista)

### Template:
- Czcionka minimum 24pt
- Spójny styl (kolory, fonty)
- Logo/branding (opcjonalnie)
- Numeracja slajdów

### Rezultat:
- Plik `Prezentacja_Zguba_gov.pptx` (lub .key, .pdf)
- Eksport do PDF jako backup

---

## Dzień 4: Nagranie wideo demo (2-3 min)

### Przygotowanie:
- [ ] Zamknąć wszystkie niepotrzebne aplikacje
- [ ] Czysty desktop (ukryj ikony)
- [ ] Wyłączyć powiadomienia (Do Not Disturb)
- [ ] Sprawdzić rozdzielczość ekranu (1920x1080 recommended)
- [ ] Przygotować przykładowe dane testowe
- [ ] Napisać scenariusz (punkt po punkcie)

### Scenariusz wideo (dokładny):

**00:00-00:10** Intro
- Ekran startowy systemu (http://localhost:4200)
- Tytuł głosem: "System Zguba.gov - prototyp jednego okna"

**00:10-01:20** Proces 5-krokowy
1. (20 sek) Wybór Warszawa z listy
2. (15 sek) Wypełnienie: "Portfel skórzany", kategoria "dokumenty", data "2025-12-07", lokalizacja "Park Łazienkowski"
3. (15 sek) Odbiór: 30 dni, "Urząd Miasta, ul. Senatorska 30", godziny "9-17", kontakt
4. (15 sek) Wybór tagów: "dokumenty", "portfele"
5. (15 sek) Klik "Opublikuj" → potwierdzenie

**01:20-01:50** Wyszukiwanie i API
- (10 sek) Wyszukanie dodanej rzeczy na liście
- (10 sek) Otwarcie Postman/Terminal
- (10 sek) Wywołanie `GET /api/found-items` → JSON response

**01:50-02:20** OData demo
- (15 sek) Wywołanie `/odata?$filter=item_status eq 'available'`
- (15 sek) Pokazanie metadanych DCAT (pola dct:)

**02:20-02:40** Responsywność
- (10 sek) Zmiana rozmiaru okna: desktop → tablet → mobile
- (10 sek) Nawigacja TAB (focus indicators)

**02:40-03:00** Statystyki i podsumowanie
- (10 sek) `/api/stats` - wyświetlenie statystyk
- (10 sek) Ekran końcowy: "Kod dostępny: github.com/..."

### Nagrywanie:
- **Narzędzie:** OBS Studio / QuickTime / Loom
- **Rozdzielczość:** 1920x1080 (Full HD)
- **Format:** MP4 (H.264)
- **Dźwięk:** Opcjonalny (może być muzyka tła + napisy)

### Post-produkcja:
- [ ] Przyciąć początek/koniec
- [ ] Dodać napisy (jeśli potrzebne)
- [ ] Dodać strzałki/highlighty (opcjonalnie)
- [ ] Sprawdzić czas (max 3 min)
- [ ] Eksport w wysokiej jakości

### Rezultat:
- Plik `Demo_Zguba_gov.mp4` (max 100 MB)
- Upload do YouTube (unlisted) lub Vimeo

---

## Dzień 5: Weryfikacja techniczna

### Backend:
- [ ] API działa: `curl http://localhost:8000/health`
- [ ] Dokumentacja Swagger dostępna: `http://localhost:8000/docs`
- [ ] Wszystkie endpointy zwracają dane
- [ ] OData endpoint działa: `/odata`
- [ ] Statystyki działają: `/api/stats`
- [ ] Baza ma przykładowe dane (min. 10 rekordów)

### Frontend:
- [ ] Uruchamia się bez błędów
- [ ] Formularz 5-krokowy działa
- [ ] Wyszukiwanie działa
- [ ] Filtry działają
- [ ] Widoki responsywne (mobile/tablet/desktop)
- [ ] Brak błędów w konsoli przeglądarki

### Testy dostępności:
- [ ] WAVE scan - brak critical errors
- [ ] Kontrast kolorów minimum AA
- [ ] Nawigacja TAB - focus visible
- [ ] Screen reader test (opcjonalnie)

### Kod:
- [ ] Repository publiczne (GitHub)
- [ ] README z instrukcją instalacji
- [ ] Wszystkie pliki skomentowane
- [ ] Brak wrażliwych danych (.env, hasła)
- [ ] Licencja (MIT/Apache)

### Rezultat:
- Działający system gotowy do demo
- Link do repository: `https://github.com/user/zguba-gov`

---

## Dzień 6: Finalizacja dokumentacji

### Sprawdzenie plików:
- [x] `README.md` - główny
- [x] `PLAN_PREZENTACJI.md` - plan slajdów
- [x] `CHECKLIST_PREZENTACJI.md` - lista kontrolna
- [x] `PRZEWODNIK_DLA_JUROROW.md` - quick start
- [x] `PRZYKLADOWE_ZAPYTANIA_API.md` - przykłady curl
- [x] `example_item_schema.json` - wzorcowy schemat
- [x] `backend/API_DOCS.md` - dokumentacja API
- [x] `backend/README.md` - backend setup
- [x] `frontend/README.md` - frontend setup

### Weryfikacja schématu danych:
- [ ] Walidacja przez JSON Schema Validator
- [ ] Min. 2 przykłady w pliku
- [ ] Wszystkie required fields
- [ ] Zgodność z DCAT-AP
- [ ] Opis każdego pola

### Rezultat:
- Kompletna dokumentacja
- Wszystkie pliki zaktualizowane
- Brak dead links

---

## Dzień 7: Generalny przegląd i backup

### Ostateczny test:
- [ ] Uruchomić system od zera (czysta instalacja)
- [ ] Przejść przez cały proces (5 kroków)
- [ ] Sprawdzić wszystkie linki w README
- [ ] Otworzyć każdy endpoint API
- [ ] Sprawdzić responsywność na prawdziwych urządzeniach
- [ ] Przejrzeć prezentację (błędy ortograficzne)
- [ ] Obejrzeć wideo (jakość, czas)

### Przygotowanie backupów:
- [ ] Kopia prezentacji (PDF + PPTX)
- [ ] Kopia wideo (lokalnie + cloud)
- [ ] Backup bazy z danymi testowymi
- [ ] Screenshot każdego slajdu (jako fallback)
- [ ] Offline copy całego projektu

### Plan B (na wypadek awarii):
- [ ] Laptop naładowany
- [ ] System działa lokalnie (bez internetu)
- [ ] Wszystkie pliki na USB drive
- [ ] Screenshots/wideo jako backup
- [ ] Notatki wydrukowane

### Rezultat:
- Gotowość 100%
- Plan awaryjny przygotowany
- Spokój przed prezentacją 😊

---

## 📋 Finalna lista dostarczalnych

### 1. Prezentacja
- [ ] `Prezentacja_Zguba_gov.pdf` (główny plik)
- [ ] `Prezentacja_Zguba_gov.pptx` (edytowalna wersja)
- [ ] Folder `/screenshots` z wszystkimi zrzutami

### 2. Prototyp
- [ ] Link do działającego demo (opcjonalnie)
- [ ] Instrukcja uruchomienia lokalnego (README.md)
- [ ] Wideo demo: `Demo_Zguba_gov.mp4`

### 3. Wzorcowy zakres danych
- [ ] `example_item_schema.json`
- [ ] Przykładowy rekord: `backend/example_item.json`

### 4. Kod źródłowy
- [ ] Repository GitHub: `https://github.com/user/zguba-gov`
- [ ] Wszystkie pliki w repo (backend + frontend)
- [ ] Dokumentacja (README, API_DOCS, etc.)

---

## 📊 Checkpoint codziennie

### Przed snem zadaj sobie pytania:

**Dzień 1-2:**
- Czy mam wszystkie potrzebne zrzuty ekranu?
- Czy screenshoty są dobrej jakości (min. 1920x1080)?

**Dzień 3:**
- Czy prezentacja ma min. 5 slajdów?
- Czy jest spójna wizualnie?
- Czy nie ma błędów ortograficznych?

**Dzień 4:**
- Czy wideo trwa 2-3 minuty?
- Czy pokazuję wszystkie kluczowe funkcje?
- Czy obraz i dźwięk są dobrej jakości?

**Dzień 5:**
- Czy system działa od A do Z?
- Czy mogę go uruchomić na czystej maszynie?

**Dzień 6:**
- Czy dokumentacja jest kompletna?
- Czy ktoś inny mógłby uruchomić projekt na podstawie README?

**Dzień 7:**
- Czy jestem gotowy na prezentację?
- Czy mam plan B?

---

## ⏰ Harmonogram dzień prezentacji

### -2 godziny przed:
- [ ] Sprawdzić laptop (bateria, porty)
- [ ] Sprawdzić internet (jeśli demo online)
- [ ] Uruchomić system lokalnie
- [ ] Otworzyć prezentację
- [ ] Mieć gotowe wideo
- [ ] Notatki do prezentacji

### -30 minut przed:
- [ ] Być na miejscu
- [ ] Sprawdzić projektor/monitor
- [ ] Test dźwięku (jeśli wideo z audio)
- [ ] Zamknąć niepotrzebne aplikacje
- [ ] Wyłączyć powiadomienia

### Podczas prezentacji:
1. Przedstawienie (1 min)
2. Slajdy (5-7 min)
3. Wideo demo (2-3 min)
4. Live demo (2-3 min) - opcjonalnie
5. Q&A (5 min)

### Po prezentacji:
- [ ] Udostępnić link do repo
- [ ] Zostawić kontakt do zespołu
- [ ] Podziękować jurorom

---

## 🎯 Kluczowe komunikaty do przekazania

### 3 najważniejsze punkty:

1. **"5 kroków i gotowe"** 
   - Prosty proces dla urzędnika bez szkolenia

2. **"Pełna zgodność z dane.gov.pl"**
   - OData + DCAT = gotowe do integracji

3. **"JSON zamiast PDF"**
   - Format odczytywalny maszynowo

### Elevator pitch (30 sekund):

> "Zguba.gov to prototyp 'jednego okna' dla rzeczy znalezionych. Urzędnik w 5 prostych krokach publikuje dane, które automatycznie udostępniamy przez API zgodne z standardem dane.gov.pl. Format JSON - odczytywalny maszynowo, nie PDF. Pełna responsywność i dostępność WCAG 2.1. Kod open source, gotowy do wdrożenia."

---

## 📞 Kontakty

**W razie pytań/problemów:**
- Zespół techniczny: _______________
- Email: _______________
- GitHub: _______________

---

## 💪 Motywacja

> "Dobra prezentacja = przygotowanie + pasja + praktyka"

- ✅ Masz świetny produkt
- ✅ Masz kompletną dokumentację
- ✅ Masz konkretny plan
- 🎯 **Teraz czas pokazać to światu!**

---

**Powodzenia!** 🚀

---

**Data utworzenia:** 7 grudnia 2025  
**Deadline prezentacji:** _______________  
**Status:** Plan gotowy - czas działać! ✅
