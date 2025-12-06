// territorial-units.service.ts
import { Injectable } from '@angular/core';

export interface TerritorialUnit {
  name: string;
  type: 'wojewodztwo' | 'powiat' | 'gmina' | 'miasto';
  parentName?: string; // np. dla powiatu: nazwa województwa
  fullName: string; // pełna nazwa z kontekstem
  id?: string; // ID z API GUS
}

interface GusApiResponse {
  totalRecords: number;
  page: number;
  pageSize: number;
  results: GusUnit[];
  links?: {
    next?: string;
  };
}

interface GusUnit {
  id: string;
  name: string;
  parentId: string;
  level: number;
  kind?: string; // "1" = powiat, "2" = miasto na prawach powiatu
}

@Injectable({
  providedIn: 'root'
})
export class TerritorialUnitsService {
  private units: TerritorialUnit[] = [];
  private wojewodztwaMap = new Map<string, string>(); // id -> nazwa
  private loading = false;
  private loaded = false;

  private readonly API_BASE = 'https://bdl.stat.gov.pl/api/v1';

  constructor() {
    this.loadDataFromApi();
  }

  private async loadDataFromApi(): Promise<void> {
    if (this.loading || this.loaded) return;
    
    this.loading = true;
    
    try {
      // 1. Pobierz województwa (level=2)
      await this.loadWojewodztwa();
      
      // 2. Pobierz powiaty i miasta na prawach powiatu (level=5)
      await this.loadPowiatyIMiasta();
      
      // 3. Pobierz gminy (level=6) - ograniczone do najważniejszych
      await this.loadGminy();
      
      this.loaded = true;
      console.log(`✅ Załadowano ${this.units.length} jednostek terytorialnych z API GUS`);
    } catch (error) {
      console.error('❌ Błąd podczas ładowania danych z API GUS:', error);
      // W razie błędu załaduj dane backup
      this.loadBackupData();
    } finally {
      this.loading = false;
    }
  }

  private async loadWojewodztwa(): Promise<void> {
    const url = `${this.API_BASE}/units?level=2&format=json&page-size=100`;
    const response = await fetch(url);
    const data: GusApiResponse = await response.json();
    
    data.results.forEach(woj => {
      const nazwa = this.formatWojewodztwoName(woj.name);
      this.wojewodztwaMap.set(woj.id, nazwa);
      
      this.units.push({
        id: woj.id,
        name: `Województwo ${nazwa}`,
        type: 'wojewodztwo',
        fullName: `Województwo ${nazwa}`
      });
    });
  }

  private async loadPowiatyIMiasta(): Promise<void> {
    let page = 0;
    let hasMore = true;
    const pageSize = 100;
    
    while (hasMore) {
      const url = `${this.API_BASE}/units?level=5&format=json&page-size=${pageSize}&page=${page}`;
      const response = await fetch(url);
      const data: GusApiResponse = await response.json();
      
      data.results.forEach(unit => {
        const wojewodztwoNazwa = this.getWojewodztwoFromParentId(unit.parentId);
        const isMiasto = unit.kind === '2';
        
        // Formatuj nazwę
        let nazwa = unit.name;
        if (isMiasto) {
          // Usuń "Powiat m. " dla miast
          nazwa = nazwa.replace(/^Powiat m\.\s*/i, '');
        }
        
        this.units.push({
          id: unit.id,
          name: nazwa,
          type: isMiasto ? 'miasto' : 'powiat',
          parentName: wojewodztwoNazwa,
          fullName: `${nazwa} (woj. ${wojewodztwoNazwa})`
        });
      });
      
      page++;
      hasMore = data.links?.next !== undefined && page < 10; // Limit dla bezpieczeństwa
    }
  }

  private async loadGminy(): Promise<void> {
    // Ładujemy tylko pierwszą stronę gmin (dla wydajności)
    // W pełnej wersji można załadować wszystkie strony
    const pageSize = 500;
    const url = `${this.API_BASE}/units?level=6&format=json&page-size=${pageSize}&page=0`;
    
    try {
      const response = await fetch(url);
      const data: GusApiResponse = await response.json();
      
      data.results.forEach(unit => {
        // Pomijamy część szczegółowych podziałów gmin
        if (unit.kind === '4' || unit.kind === '5') {
          return; // Pomiń miasta i obszary wiejskie w gminach miejsko-wiejskich
        }
        
        const wojewodztwoNazwa = this.getWojewodztwoFromUnitId(unit.id);
        
        // Formatuj nazwę
        let nazwa = unit.name;
        if (!nazwa.toLowerCase().startsWith('gmina')) {
          nazwa = `Gmina ${nazwa}`;
        }
        
        this.units.push({
          id: unit.id,
          name: nazwa,
          type: 'gmina',
          parentName: wojewodztwoNazwa,
          fullName: `${nazwa} (woj. ${wojewodztwoNazwa})`
        });
      });
    } catch (error) {
      console.warn('Nie udało się załadować gmin:', error);
    }
  }

  private getWojewodztwoFromParentId(parentId: string): string {
    // parentId podregionu zawiera kod województwa
    const wojCode = parentId.substring(0, 4) + '00000000';
    return this.wojewodztwaMap.get(wojCode) || 'nieznane';
  }

  private getWojewodztwoFromUnitId(unitId: string): string {
    // ID jednostki zawiera kod województwa w pierwszych cyfrach
    const wojCode = unitId.substring(0, 4) + '00000000';
    return this.wojewodztwaMap.get(wojCode) || 'nieznane';
  }

  private formatWojewodztwoName(name: string): string {
    // API zwraca nazwy wielkimi literami, formatujemy je
    return name.toLowerCase()
      .split(' ')
      .map(word => word.charAt(0).toUpperCase() + word.slice(1))
      .join(' ')
      .replace(/-/g, '-');
  }

  private loadBackupData(): void {
    // Województwa
    const wojewodztwa = [
      'dolnośląskie', 'kujawsko-pomorskie', 'lubelskie', 'lubuskie',
      'łódzkie', 'małopolskie', 'mazowieckie', 'opolskie',
      'podkarpackie', 'podlaskie', 'pomorskie', 'śląskie',
      'świętokrzyskie', 'warmińsko-mazurskie', 'wielkopolskie', 'zachodniopomorskie'
    ];

    wojewodztwa.forEach(woj => {
      this.units.push({
        name: `Województwo ${woj}`,
        type: 'wojewodztwo',
        fullName: `Województwo ${woj}`
      });
    });

    // Przykładowe powiaty (reprezentatywna lista najważniejszych)
    const powiaty = [
      { name: 'Powiat warszawski zachodni', wojewodztwo: 'mazowieckie' },
      { name: 'Powiat piaseczyński', wojewodztwo: 'mazowieckie' },
      { name: 'Powiat pruszkowski', wojewodztwo: 'mazowieckie' },
      { name: 'Powiat wołomiński', wojewodztwo: 'mazowieckie' },
      { name: 'Powiat grodziski', wojewodztwo: 'mazowieckie' },
      { name: 'Powiat krakowski', wojewodztwo: 'małopolskie' },
      { name: 'Powiat wielicki', wojewodztwo: 'małopolskie' },
      { name: 'Powiat myślenicki', wojewodztwo: 'małopolskie' },
      { name: 'Powiat poznański', wojewodztwo: 'wielkopolskie' },
      { name: 'Powiat wrocławski', wojewodztwo: 'dolnośląskie' },
      { name: 'Powiat gdański', wojewodztwo: 'pomorskie' },
      { name: 'Powiat katowicki', wojewodztwo: 'śląskie' },
      { name: 'Powiat bielski', wojewodztwo: 'śląskie' },
      { name: 'Powiat gliwicki', wojewodztwo: 'śląskie' },
      { name: 'Powiat lubliniecki', wojewodztwo: 'śląskie' },
      { name: 'Powiat łódzki wschodni', wojewodztwo: 'łódzkie' },
      { name: 'Powiat rzeszowski', wojewodztwo: 'podkarpackie' },
      { name: 'Powiat białostocki', wojewodztwo: 'podlaskie' },
      { name: 'Powiat szczeciński', wojewodztwo: 'zachodniopomorskie' },
    ];

    powiaty.forEach(powiat => {
      this.units.push({
        name: powiat.name,
        type: 'powiat',
        parentName: powiat.wojewodztwo,
        fullName: `${powiat.name} (woj. ${powiat.wojewodztwo})`
      });
    });

    // Miasta na prawach powiatu
    const miastaNaPrawachPowiatu = [
      { name: 'Warszawa', wojewodztwo: 'mazowieckie' },
      { name: 'Kraków', wojewodztwo: 'małopolskie' },
      { name: 'Wrocław', wojewodztwo: 'dolnośląskie' },
      { name: 'Poznań', wojewodztwo: 'wielkopolskie' },
      { name: 'Gdańsk', wojewodztwo: 'pomorskie' },
      { name: 'Szczecin', wojewodztwo: 'zachodniopomorskie' },
      { name: 'Bydgoszcz', wojewodztwo: 'kujawsko-pomorskie' },
      { name: 'Lublin', wojewodztwo: 'lubelskie' },
      { name: 'Katowice', wojewodztwo: 'śląskie' },
      { name: 'Białystok', wojewodztwo: 'podlaskie' },
      { name: 'Gdynia', wojewodztwo: 'pomorskie' },
      { name: 'Częstochowa', wojewodztwo: 'śląskie' },
      { name: 'Radom', wojewodztwo: 'mazowieckie' },
      { name: 'Sosnowiec', wojewodztwo: 'śląskie' },
      { name: 'Toruń', wojewodztwo: 'kujawsko-pomorskie' },
      { name: 'Kielce', wojewodztwo: 'świętokrzyskie' },
      { name: 'Rzeszów', wojewodztwo: 'podkarpackie' },
      { name: 'Gliwice', wojewodztwo: 'śląskie' },
      { name: 'Zabrze', wojewodztwo: 'śląskie' },
      { name: 'Olsztyn', wojewodztwo: 'warmińsko-mazurskie' },
      { name: 'Bielsko-Biała', wojewodztwo: 'śląskie' },
      { name: 'Bytom', wojewodztwo: 'śląskie' },
      { name: 'Zielona Góra', wojewodztwo: 'lubuskie' },
      { name: 'Rybnik', wojewodztwo: 'śląskie' },
      { name: 'Ruda Śląska', wojewodztwo: 'śląskie' },
      { name: 'Opole', wojewodztwo: 'opolskie' },
      { name: 'Tychy', wojewodztwo: 'śląskie' },
      { name: 'Gorzów Wielkopolski', wojewodztwo: 'lubuskie' },
      { name: 'Elbląg', wojewodztwo: 'warmińsko-mazurskie' },
      { name: 'Dąbrowa Górnicza', wojewodztwo: 'śląskie' },
      { name: 'Płock', wojewodztwo: 'mazowieckie' },
      { name: 'Wałbrzych', wojewodztwo: 'dolnośląskie' },
      { name: 'Włocławek', wojewodztwo: 'kujawsko-pomorskie' },
      { name: 'Tarnów', wojewodztwo: 'małopolskie' },
      { name: 'Chorzów', wojewodztwo: 'śląskie' },
      { name: 'Koszalin', wojewodztwo: 'zachodniopomorskie' },
      { name: 'Kalisz', wojewodztwo: 'wielkopolskie' },
      { name: 'Legnica', wojewodztwo: 'dolnośląskie' },
      { name: 'Grudziądz', wojewodztwo: 'kujawsko-pomorskie' },
      { name: 'Słupsk', wojewodztwo: 'pomorskie' },
      { name: 'Jaworzno', wojewodztwo: 'śląskie' },
      { name: 'Jastrzębie-Zdrój', wojewodztwo: 'śląskie' },
      { name: 'Nowy Sącz', wojewodztwo: 'małopolskie' },
      { name: 'Jelenia Góra', wojewodztwo: 'dolnośląskie' },
      { name: 'Siedlce', wojewodztwo: 'mazowieckie' },
      { name: 'Mysłowice', wojewodztwo: 'śląskie' },
      { name: 'Konin', wojewodztwo: 'wielkopolskie' },
      { name: 'Piotrków Trybunalski', wojewodztwo: 'łódzkie' },
      { name: 'Inowrocław', wojewodztwo: 'kujawsko-pomorskie' },
      { name: 'Lubin', wojewodztwo: 'dolnośląskie' },
      { name: 'Ostrów Wielkopolski', wojewodztwo: 'wielkopolskie' },
      { name: 'Suwałki', wojewodztwo: 'podlaskie' },
      { name: 'Stargard', wojewodztwo: 'zachodniopomorskie' },
      { name: 'Gniezno', wojewodztwo: 'wielkopolskie' },
      { name: 'Głogów', wojewodztwo: 'dolnośląskie' },
      { name: 'Pabianice', wojewodztwo: 'łódzkie' },
      { name: 'Chełm', wojewodztwo: 'lubelskie' },
      { name: 'Zamość', wojewodztwo: 'lubelskie' },
      { name: 'Tomaszów Mazowiecki', wojewodztwo: 'łódzkie' },
      { name: 'Przemyśl', wojewodztwo: 'podkarpackie' },
      { name: 'Stalowa Wola', wojewodztwo: 'podkarpackie' },
      { name: 'Mielec', wojewodztwo: 'podkarpackie' },
      { name: 'Krosno', wojewodztwo: 'podkarpackie' },
      { name: 'Tarnobrzeg', wojewodztwo: 'podkarpackie' },
      { name: 'Łomża', wojewodztwo: 'podlaskie' },
      { name: 'Biała Podlaska', wojewodztwo: 'lubelskie' },
      { name: 'Skierniewice', wojewodztwo: 'łódzkie' },
      { name: 'Sopot', wojewodztwo: 'pomorskie' },
      { name: 'Świnoujście', wojewodztwo: 'zachodniopomorskie' }
    ];

    miastaNaPrawachPowiatu.forEach(miasto => {
      this.units.push({
        name: miasto.name,
        type: 'miasto',
        parentName: miasto.wojewodztwo,
        fullName: `Miasto ${miasto.name} (woj. ${miasto.wojewodztwo})`
      });
    });

    // Przykładowe gminy (reprezentatywne dla największych aglomeracji)
    const gminy = [
      { name: 'Gmina Lesznowola', wojewodztwo: 'mazowieckie' },
      { name: 'Gmina Konstancin-Jeziorna', wojewodztwo: 'mazowieckie' },
      { name: 'Gmina Michałowice', wojewodztwo: 'mazowieckie' },
      { name: 'Gmina Raszyn', wojewodztwo: 'mazowieckie' },
      { name: 'Gmina Wieliczka', wojewodztwo: 'małopolskie' },
      { name: 'Gmina Zabierzów', wojewodztwo: 'małopolskie' },
      { name: 'Gmina Tarnowo Podgórne', wojewodztwo: 'wielkopolskie' },
      { name: 'Gmina Komorniki', wojewodztwo: 'wielkopolskie' },
      { name: 'Gmina Długołęka', wojewodztwo: 'dolnośląskie' },
      { name: 'Gmina Kobierzyce', wojewodztwo: 'dolnośląskie' },
      { name: 'Gmina Żukowo', wojewodztwo: 'pomorskie' },
      { name: 'Gmina Pruszcz Gdański', wojewodztwo: 'pomorskie' },
    ];

    gminy.forEach(gmina => {
      this.units.push({
        name: gmina.name,
        type: 'gmina',
        parentName: gmina.wojewodztwo,
        fullName: `${gmina.name} (woj. ${gmina.wojewodztwo})`
      });
    });
  }

  /**
   * Wyszukuje jednostki terytorialne na podstawie zapytania
   * @param query - fraza do wyszukania
   * @param type - opcjonalny filtr typu jednostki
   * @returns tablica pasujących jednostek
   */
  async search(query: string, type?: TerritorialUnit['type']): Promise<TerritorialUnit[]> {
    // Czekaj na załadowanie danych jeśli jeszcze nie są gotowe
    if (this.loading) {
      await this.waitForLoad();
    }
    
    if (!query || query.length < 2) {
      return [];
    }

    const normalizedQuery = this.normalize(query);
    
    let filtered = this.units.filter(unit => {
      const matchesQuery = this.normalize(unit.fullName).includes(normalizedQuery) ||
                           this.normalize(unit.name).includes(normalizedQuery);
      const matchesType = !type || unit.type === type;
      return matchesQuery && matchesType;
    });

    // Sortuj: najpierw te które zaczynają się od query, potem reszta
    filtered.sort((a, b) => {
      const aStarts = this.normalize(a.name).startsWith(normalizedQuery);
      const bStarts = this.normalize(b.name).startsWith(normalizedQuery);
      
      if (aStarts && !bStarts) return -1;
      if (!aStarts && bStarts) return 1;
      return a.name.localeCompare(b.name, 'pl');
    });

    return filtered.slice(0, 20); // Zwróć maksymalnie 20 wyników
  }

  /**
   * Czeka na załadowanie danych
   */
  private waitForLoad(): Promise<void> {
    return new Promise((resolve) => {
      const checkInterval = setInterval(() => {
        if (!this.loading || this.loaded) {
          clearInterval(checkInterval);
          resolve();
        }
      }, 100);
      
      // Timeout po 10 sekundach
      setTimeout(() => {
        clearInterval(checkInterval);
        resolve();
      }, 10000);
    });
  }

  /**
   * Pobiera wszystkie jednostki określonego typu
   */
  getByType(type: TerritorialUnit['type']): TerritorialUnit[] {
    return this.units.filter(unit => unit.type === type);
  }

  /**
   * Normalizuje tekst dla porównania (małe litery, bez polskich znaków)
   */
  private normalize(text: string): string {
    return text
      .toLowerCase()
      .replace(/ą/g, 'a')
      .replace(/ć/g, 'c')
      .replace(/ę/g, 'e')
      .replace(/ł/g, 'l')
      .replace(/ń/g, 'n')
      .replace(/ó/g, 'o')
      .replace(/ś/g, 's')
      .replace(/ź/g, 'z')
      .replace(/ż/g, 'z');
  }
}
