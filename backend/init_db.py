#!/usr/bin/env python3
"""
Initialize database - creates tables for found items
"""
import asyncio
from database import init_db


async def main():
    """Main initialization function"""
    print("🚀 Inicjalizacja bazy danych...")
    print("=" * 60)
    
    # Create tables
    print("📋 Tworzenie tabel...")
    await init_db()
    print("✅ Tabele utworzone!")
    
    print("\n" + "=" * 60)
    print("✅ Inicjalizacja zakończona pomyślnie!")
    print("\nMożesz teraz uruchomić serwer:")
    print("  uvicorn main:app --reload")
    print("  lub: bash start.sh")


if __name__ == "__main__":
    asyncio.run(main())
