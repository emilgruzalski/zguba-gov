#!/usr/bin/env python3
"""
Add example found items to database for testing
"""
import asyncio
import uuid
from datetime import datetime
from database import AsyncSessionLocal
from models import FoundItem


async def add_example_items():
    """Add example found items"""
    
    example_items = [
        {
            "municipality_name": "Warszawa",
            "municipality_type": "miasto",
            "municipality_email": "kontakt@um.warszawa.pl",
            "item_name": "Portfel skórzany brązowy",
            "item_category": "dokumenty",
            "item_date": "2025-12-01",
            "item_location": "Park Łazienkowski, ławka przy fontannie",
            "item_status": "available",
            "item_description": "Brązowy portfel ze skóry naturalnej, zawiera dowód osobisty i prawo jazdy",
            "pickup_deadline": 30,
            "pickup_location": "Urząd Miasta Warszawa, ul. Senatorska 30, pokój 215",
            "pickup_hours": "Poniedziałek-Piątek 9:00-17:00",
            "pickup_contact": "+48 22 443 44 44",
            "categories": ["dokumenty", "portfele", "dowód osobisty"]
        },
        {
            "municipality_name": "Kraków",
            "municipality_type": "miasto",
            "municipality_email": "kontakt@um.krakow.pl",
            "item_name": "iPhone 14 Pro",
            "item_category": "telefony",
            "item_date": "2025-12-02",
            "item_location": "Rynek Główny, przy fontannie",
            "item_status": "available",
            "item_description": "Czarny iPhone 14 Pro, ekran lekko zarysowany",
            "pickup_deadline": 30,
            "pickup_location": "Urząd Miasta Krakowa, ul. Wszystkich Świętych 3",
            "pickup_hours": "Poniedziałek-Piątek 8:00-16:00",
            "pickup_contact": "+48 12 616 15 00",
            "categories": ["telefony", "elektronika"]
        },
        {
            "municipality_name": "Gdańsk",
            "municipality_type": "miasto",
            "municipality_email": "biuro@gdansk.pl",
            "item_name": "Plecak turystyczny niebieski",
            "item_category": "bagaże",
            "item_date": "2025-12-03",
            "item_location": "Dworzec PKP Gdańsk Główny",
            "item_status": "available",
            "item_description": "Niebieski plecak turystyczny marki Deuter, zawiera rzeczy osobiste",
            "pickup_deadline": 30,
            "pickup_location": "Urząd Miasta Gdańska, ul. Nowe Ogrody 8/12",
            "pickup_hours": "Poniedziałek-Piątek 7:30-15:30",
            "pickup_contact": "+48 58 323 66 00",
            "categories": ["bagaże", "plecaki"]
        },
        {
            "municipality_name": "Wrocław",
            "municipality_type": "miasto",
            "municipality_email": "um@um.wroc.pl",
            "item_name": "Klucze z breloczkiem z logiem VW",
            "item_category": "klucze",
            "item_date": "2025-12-04",
            "item_location": "Plac Solny, przy kwiaciarniach",
            "item_status": "available",
            "item_description": "Zestaw kluczy z pilotem do samochodu Volkswagen",
            "pickup_deadline": 30,
            "pickup_location": "Urząd Miejski Wrocławia, pl. Nowy Targ 1-8",
            "pickup_hours": "Poniedziałek 8:00-18:00, Wtorek-Piątek 7:30-15:30",
            "pickup_contact": "+48 71 777 77 77",
            "categories": ["klucze", "pilot samochodowy"]
        },
        {
            "municipality_name": "Poznań",
            "municipality_type": "miasto",
            "municipality_email": "urzad@um.poznan.pl",
            "item_name": "Okulary przeciwsłoneczne Ray-Ban",
            "item_category": "odzież",
            "item_date": "2025-12-05",
            "item_location": "Stary Rynek, przy ratuszu",
            "item_status": "available",
            "item_description": "Czarne okulary Ray-Ban Wayfarer w etui",
            "pickup_deadline": 30,
            "pickup_location": "Urząd Miasta Poznania, pl. Kolegiacki 17",
            "pickup_hours": "Poniedziałek-Piątek 8:00-16:00",
            "pickup_contact": "+48 61 878 40 00",
            "categories": ["odzież", "akcesoria", "okulary"]
        }
    ]
    
    async with AsyncSessionLocal() as session:
        try:
            print("📥 Dodawanie przykładowych danych...")
            
            for item_data in example_items:
                item = FoundItem(
                    id=str(uuid.uuid4()),
                    **item_data
                )
                session.add(item)
                print(f"   ✓ {item_data['item_name']} - {item_data['municipality_name']}")
            
            await session.commit()
            print(f"\n✅ Dodano {len(example_items)} przykładowych przedmiotów!")
            
        except Exception as e:
            print(f"❌ Błąd: {e}")
            await session.rollback()


async def main():
    print("🚀 Dodawanie przykładowych danych do bazy...")
    print("=" * 60)
    await add_example_items()
    print("=" * 60)
    print("\n💡 Możesz teraz sprawdzić dane:")
    print("   curl http://localhost:8000/api/found-items")
    print("   lub odwiedź: http://localhost:8000/docs")


if __name__ == "__main__":
    asyncio.run(main())
