import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

export interface MunicipalityInfo {
  name: string;
  type: string;
  contactEmail: string;
}

export interface ItemInfo {
  name: string;
  category: string;
  date: string;
  location: string;
  status: string;
  description?: string;
}

export interface PickupInfo {
  deadline: number;
  location: string;
  hours?: string;
  contact?: string;
}

export interface FoundItemCreate {
  municipality: MunicipalityInfo;
  item: ItemInfo;
  pickup: PickupInfo;
  categories?: string[];
}

export interface FoundItemResponse extends FoundItemCreate {
  id: string;
  createdAt: string;
  updatedAt?: string;
}

@Injectable({ providedIn: 'root' })
export class FoundItemsService {
  private http = inject(HttpClient);

  private readonly apiUrl = environment.apiUrl;

  list(skip = 0, limit = 50, params: Record<string, string> = {}): Observable<FoundItemResponse[]> {
    let url = `${this.apiUrl}/found-items?skip=${skip}&limit=${limit}`;
    Object.keys(params).forEach(key => {
      if (params[key] !== undefined && params[key] !== null) {
        url += `&${encodeURIComponent(key)}=${encodeURIComponent(params[key])}`;
      }
    });
    return this.http.get<FoundItemResponse[]>(url);
  }

  create(data: FoundItemCreate): Observable<FoundItemResponse> {
    return this.http.post<FoundItemResponse>(`${this.apiUrl}/found-items`, data);
  }

  get(id: string): Observable<FoundItemResponse> {
    return this.http.get<FoundItemResponse>(`${this.apiUrl}/found-items/${encodeURIComponent(id)}`);
  }

  update(id: string, data: Partial<FoundItemCreate>): Observable<FoundItemResponse> {
    return this.http.put<FoundItemResponse>(`${this.apiUrl}/found-items/${encodeURIComponent(id)}`, data);
  }

  delete(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/found-items/${encodeURIComponent(id)}`);
  }

  categoriesList(): Observable<{value:string;label:string}[]> {
    return this.http.get<{value:string;label:string}[]>(`${this.apiUrl}/found-items/categories/list`);
  }
}
