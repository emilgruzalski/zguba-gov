// app.component.ts
import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormBuilder, FormGroup, Validators } from '@angular/forms';

interface FoundItem {
  id: string;
  municipality: {
    name: string;
    type: string;
    contactEmail: string;
  };
  item: {
    name: string;
    category: string;
    date: string;
    location: string;
    status: string;
    description?: string;
  };
  pickup: {
    deadline: number;
    location: string;
    hours?: string;
    contact?: string;
  };
  categories: string[];
  createdAt: string;
}

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  currentStep: number = 1;
  totalSteps: number = 5;
  
  form: FormGroup;
  items: FoundItem[] = [];
  
  categories = [
    { value: 'dokumenty', label: 'Dokumenty' },
    { value: 'telefony', label: 'Telefony i elektronika' },
    { value: 'odziez', label: 'Odzież i akcesoria' },
    { value: 'kluczyki', label: 'Klucze' },
    { value: 'bizesmeny', label: 'Torby i plecaki' },
    { value: 'inne', label: 'Inne' }
  ];

  municipalityTypes = [
    { value: 'powiat', label: 'Powiat' },
    { value: 'gmina', label: 'Gmina' },
    { value: 'miasto', label: 'Miasto' },
    { value: 'wojewodztwo', label: 'Województwo' }
  ];

  statusOptions = [
    { value: 'oczekuje', label: 'Oczekuje na właściciela' },
    { value: 'odebrana', label: 'Odebrana przez właściciela' },
    { value: 'przekazana', label: 'Przekazana organizacji' }
  ];

  constructor(private fb: FormBuilder) {
    this.form = this.initializeForm();
  }

  ngOnInit(): void {
    this.loadItems();
  }

  initializeForm(): FormGroup {
    return this.fb.group({
      // Krok 1
      municipalityName: ['', [Validators.required]],
      municipalityType: ['', [Validators.required]],
      contactEmail: ['', [Validators.required, Validators.email]],
      
      // Krok 2
      categories: [[], [Validators.required]],
      
      // Krok 3
      itemName: ['', [Validators.required]],
      itemCategory: ['', [Validators.required]],
      itemDate: ['', [Validators.required]],
      itemLocation: ['', [Validators.required]],
      itemStatus: ['', [Validators.required]],
      itemDescription: [''],
      
      // Krok 4
      storageDeadline: [30, [Validators.required, Validators.min(1), Validators.max(365)]],
      pickupLocation: ['', [Validators.required]],
      pickupHours: [''],
      contactPerson: ['']
    });
  }

  nextStep(): void {
    if (this.validateStep(this.currentStep)) {
      this.currentStep++;
    }
  }

  prevStep(): void {
    if (this.currentStep > 1) {
      this.currentStep--;
    }
  }

  validateStep(step: number): boolean {
    const controls = this.form.controls;
    
    switch (step) {
      case 1:
        return controls['municipalityName'].valid && 
               controls['municipalityType'].valid && 
               controls['contactEmail'].valid;
      case 2:
        return controls['categories'].value.length > 0;
      case 3:
        return controls['itemName'].valid && 
               controls['itemCategory'].valid && 
               controls['itemDate'].valid && 
               controls['itemLocation'].valid && 
               controls['itemStatus'].valid;
      case 4:
        return controls['storageDeadline'].valid && 
               controls['pickupLocation'].valid;
      default:
        return true;
    }
  }

  submitData(): void {
    const newItem: FoundItem = {
      id: this.generateId(),
      municipality: {
        name: this.form.value.municipalityName,
        type: this.form.value.municipalityType,
        contactEmail: this.form.value.contactEmail
      },
      item: {
        name: this.form.value.itemName,
        category: this.form.value.itemCategory,
        date: this.form.value.itemDate,
        location: this.form.value.itemLocation,
        status: this.form.value.itemStatus,
        description: this.form.value.itemDescription
      },
      pickup: {
        deadline: this.form.value.storageDeadline,
        location: this.form.value.pickupLocation,
        hours: this.form.value.pickupHours,
        contact: this.form.value.contactPerson
      },
      categories: this.form.value.categories,
      createdAt: new Date().toISOString()
    };

    this.items.push(newItem);
    this.saveItems();
    alert('✓ Dane zostały pomyślnie udostępnione!');
    this.resetForm();
    this.currentStep = 1;
  }

  exportJSON(): void {
    const dataStr = JSON.stringify(this.items, null, 2);
    this.downloadFile(dataStr, `rzeczy_znalezione_${Date.now()}.json`, 'application/json');
  }

  exportCSV(): void {
    let csv = 'ID,Samorząd,Typ,Email,Przedmiot,Kategoria,Data,Lokalizacja,Status,Odbiór,Godziny,Data dodania\n';
    
    this.items.forEach(item => {
      const row = [
        item.id,
        item.municipality.name,
        item.municipality.type,
        item.municipality.contactEmail,
        item.item.name,
        item.item.category,
        item.item.date,
        item.item.location,
        item.item.status,
        item.pickup.location,
        item.pickup.hours || '',
        item.createdAt
      ];
      csv += row.map(cell => `"${cell}"`).join(',') + '\n';
    });

    this.downloadFile(csv, `rzeczy_znalezione_${Date.now()}.csv`, 'text/csv');
  }

  private downloadFile(content: string, filename: string, type: string): void {
    const element = document.createElement('a');
    element.setAttribute('href', `data:${type};charset=utf-8,${encodeURIComponent(content)}`);
    element.setAttribute('download', filename);
    element.style.display = 'none';
    document.body.appendChild(element);
    element.click();
    document.body.removeChild(element);
  }

  private generateId(): string {
    return 'RZ-' + Date.now() + '-' + Math.random().toString(36).substr(2, 9);
  }

  private saveItems(): void {
    localStorage.setItem('foundItems', JSON.stringify(this.items));
  }

  private loadItems(): void {
    const stored = localStorage.getItem('foundItems');
    if (stored) {
      this.items = JSON.parse(stored);
    }
  }

  private resetForm(): void {
    this.form.reset({
      storageDeadline: 30
    });
  }

  toggleCategory(value: string): void {
    const categoriesArray = this.form.get('categories') as any;
    const index = categoriesArray.value.indexOf(value);
    
    if (index === -1) {
      categoriesArray.value.push(value);
    } else {
      categoriesArray.value.splice(index, 1);
    }
  }

  getProgress(): number {
    return (this.currentStep / this.totalSteps) * 100;
  }

  isStepCompleted(step: number): boolean {
    return step < this.currentStep;
  }

  isStepActive(step: number): boolean {
    return step === this.currentStep;
  }
}