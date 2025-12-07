# 🚀 Quick Start - Zguba.gov

## Dla spieszczących się (2 minuty)

### 1️⃣ Uruchom system
```bash
# Terminal 1: Backend
cd backend
python3 -m venv venv && source venv/bin/activate
pip install -r requirements.txt
python init_db.py && python add_examples.py
uvicorn main:app --reload

# Terminal 2: Frontend
cd frontend
npm install && npm start
```

### 2️⃣ Otwórz przeglądarki
- Frontend: http://localhost:4200
- API Docs: http://localhost:8000/docs
- OData: http://localhost:8000/odata

### 3️⃣ Testuj 5-krokowy proces
1. Wybierz gminę (np. "Warszawa")
2. Wpisz dane rzeczy (portfel, dokumenty, 2025-12-07, park)
3. Miejsce odbioru (urząd, 30 dni)
4. Wybierz tagi (dokumenty, portfele)
5. Kliknij "Opublikuj"

---

## 📚 Dokumentacja

| Plik | Opis | Dla kogo |
|------|------|----------|
| **[PRZEWODNIK_DLA_JUROROW.md](PRZEWODNIK_DLA_JUROROW.md)** | Start 5 min, weryfikacja wymagań | ⭐ Jurorzy |
| **[PLAN_PREZENTACJI.md](PLAN_PREZENTACJI.md)** | 8 slajdów, scenariusz wideo | ⭐ Prezentujący |
| **[CHECKLIST_PREZENTACJI.md](CHECKLIST_PREZENTACJI.md)** | Co zrobić przed prezentacją | Prezentujący |
| **[HARMONOGRAM_PREZENTACJI.md](HARMONOGRAM_PREZENTACJI.md)** | Plan 7-dniowy | Prezentujący |
| **[example_item_schema.json](example_item_schema.json)** | Wzorcowy format danych | ⭐ Wszyscy |
| **[PRZYKLADOWE_ZAPYTANIA_API.md](PRZYKLADOWE_ZAPYTANIA_API.md)** | Przykłady curl, Postman | Deweloperzy |

---

## ✅ Spełnienie wymagań

### Wymagania formalne ✓
- ✅ Prezentacja min. 5 slajdów → Plan na 8 slajdów
- ✅ Zrzuty ekranu z dowodami → Lista w planie
- ✅ Opis funkcjonalności → README + Plan
- ✅ Opis kroków użytkownika → 5 kroków opisanych
- ✅ Lista ograniczeń → W planie (Slajd 8)
- ✅ Działający prototyp → Backend + Frontend
- ✅ Wideo 2-3 min → Scenariusz gotowy
- ✅ Wzorcowy zakres danych → example_item_schema.json
- ✅ Kod źródłowy → Cały projekt

### Wymagania techniczne ✓
- ✅ UX - max 5 kroków
- ✅ Format maszynoczytelny (JSON, nie PDF)
- ✅ Zgodność z dane.gov.pl (OData + DCAT)
- ✅ WCAG 2.1 (kontrast, ARIA, klawiatura)
- ✅ Responsywność (mobile/tablet/desktop)

---

## 🎯 Kluczowe endpointy

```bash
# Health check
curl http://localhost:8000/health

# Lista rzeczy
curl http://localhost:8000/api/found-items

# OData (dane.gov.pl)
curl http://localhost:8000/odata

# Statystyki
curl http://localhost:8000/api/stats
```

---

## 📞 Problemy?

**System nie startuje?**
```bash
# Sprawdź Python
python3 --version  # Minimum 3.9

# Sprawdź Node
node --version     # Minimum 18

# Reinstaluj
cd backend && pip install -r requirements.txt
cd frontend && npm install
```

**API nie zwraca danych?**
```bash
cd backend
rm zguba.db
python init_db.py
python add_examples.py
```

---

## 💡 Najważniejsze 3 punkty

1. **5 kroków** - prosty proces dla urzędnika
2. **OData + DCAT** - zgodność z dane.gov.pl
3. **JSON** - odczytywalny maszynowo (nie PDF)

---

**Start:** [PRZEWODNIK_DLA_JUROROW.md](PRZEWODNIK_DLA_JUROROW.md)  
**Prezentacja:** [PLAN_PREZENTACJI.md](PLAN_PREZENTACJI.md)  
**Schemat danych:** [example_item_schema.json](example_item_schema.json)
