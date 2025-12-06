const XLSX = require('xlsx');
const fs = require('fs');
const path = require('path');

// Funkcja normalizująca polskie znaki
const normalize = (text) => {
  if (!text) return '';
  return text.toString().toLowerCase()
    .replace(/ą/g, 'a').replace(/ć/g, 'c').replace(/ę/g, 'e')
    .replace(/ł/g, 'l').replace(/ń/g, 'n').replace(/ó/g, 'o')
    .replace(/ś/g, 's').replace(/ź/g, 'z').replace(/ż/g, 'z');
};

// Ścieżka do pliku Excel
const excelPath = path.join(__dirname, '../db/baza_jst_aktualizacja_26062025.xlsx');
const outputPath = path.join(__dirname, '../src/assets/territorial-units.json');

console.log('Odczytuję plik Excel:', excelPath);

// Wczytaj plik Excel
const workbook = XLSX.readFile(excelPath);
const sheetName = workbook.SheetNames[0];
const worksheet = workbook.Sheets[sheetName];

// Konwertuj na JSON
const data = XLSX.utils.sheet_to_json(worksheet);

console.log('Znaleziono rekordów:', data.length);
console.log('Przykładowy rekord:', data[0]);

// Przeanalizuj strukturę danych
if (data.length > 0) {
  console.log('\nKolumny w pliku:');
  Object.keys(data[0]).forEach(key => {
    console.log(`  - ${key}`);
  });
}

// Mapuj dane do struktury używanej w aplikacji
const units = data.map((row, index) => {
  // Określ typ jednostki na podstawie typ_JST
  let type = 'gmina';
  let level = 6;
  let kind = '1';
  
  const typJST = row['typ_JST'];
  
  if (typJST === 'W') {
    type = 'wojewodztwo';
    level = 2;
  } else if (typJST === 'P') {
    type = 'powiat';
    level = 5;
    kind = '1';
  } else if (typJST === 'MNP') {
    // MNP = Miasto Na Prawach powiatu
    type = 'miasto';
    level = 5;
    kind = '2';
  } else if (typJST === 'dzielnica') {
    // Pomijamy dzielnice
    return null;
  } else {
    // GW, GM, GMW - różne typy gmin
    type = 'gmina';
    level = 6;
    
    if (typJST === 'GW') kind = '2'; // gmina wiejska
    else if (typJST === 'GM') kind = '3'; // gmina miejska
    else if (typJST === 'GMW') kind = '1'; // gmina miejsko-wiejska
    else kind = '1';
  }
  
  // Pobierz nazwę samorządu
  const name = row['nazwa_samorządu'] || '';
  
  // Pobierz email (jeśli istnieje)
  const email = row['ogólny adres poczty elektronicznej gminy/powiatu/województwa'] || '';
  
  return {
    id: row['Kod_TERYT'] || `unit-${index}`,
    name: name,
    parentId: '', // Możemy dodać logikę dla parentId jeśli potrzebne
    level: level,
    kind: kind,
    type: type,
    hasDescription: false,
    email: email, // Dodajemy email z pliku
    officeName: row['nazwa_urzędu_JST'] || '',
    voivodeship: row['Województwo'] || '',
    county: row['Powiat'] || ''
  };
}).filter(u => u !== null); // Usuń dzielnice

// Zapisz do pliku JSON
fs.mkdirSync(path.dirname(outputPath), { recursive: true });
fs.writeFileSync(outputPath, JSON.stringify(units, null, 2));

console.log('\nKonwersja zakończona!');
console.log('Plik JSON zapisany:', outputPath);
console.log('Liczba jednostek:', units.length);

// Statystyki
const stats = {
  wojewodztwa: units.filter(u => u.type === 'wojewodztwo').length,
  powiaty: units.filter(u => u.type === 'powiat').length,
  miasta: units.filter(u => u.type === 'miasto').length,
  gminy: units.filter(u => u.type === 'gmina').length
};

console.log('\nStatystyki:');
console.log('  Województwa:', stats.wojewodztwa);
console.log('  Powiaty:', stats.powiaty);
console.log('  Miasta:', stats.miasta);
console.log('  Gminy:', stats.gminy);
