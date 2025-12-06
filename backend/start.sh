#!/bin/bash

# Start Backend Script

echo "🚀 Uruchamianie backendu FastAPI..."

cd "$(dirname "$0")"

# Sprawdź czy wirtualne środowisko istnieje
if [ ! -d "venv" ]; then
    echo "❌ Wirtualne środowisko nie istnieje!"
    echo "Uruchom najpierw: python3 -m venv venv && source venv/bin/activate && pip install -r requirements.txt"
    exit 1
fi

# Sprawdź czy baza danych istnieje
if [ ! -f "zguba_gov.db" ]; then
    echo "⚠️  Baza danych nie istnieje. Inicjalizacja..."
    ./venv/bin/python init_db.py
fi

# Uruchom serwer
echo "✅ Uruchamianie serwera na http://localhost:8000"
./venv/bin/uvicorn main:app --reload --host 0.0.0.0 --port 8000
