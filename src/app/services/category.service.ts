import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Category } from '../mock-data/category';
import { CATEGORIES } from '../mock-data/mock-categories';
import { EventEmitter } from '@angular/core';

export interface User {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  firstname: string;
  lastname: string;
  email: string;
  password: string;
}

@Injectable({
  providedIn: 'root'
})
export class CategoryService {
  private apiURL = 'http://localhost:9000';

  constructor(private http: HttpClient) { }
  
  getAPI(): Observable<User[]> {
    const url = `${this.apiURL}/users`;
    return this.http.get<User[]>(url);
  }

  getCategories(): Category[] {
    return CATEGORIES;
  }
}

