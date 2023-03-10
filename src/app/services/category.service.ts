import { Injectable } from '@angular/core';
import { Category } from '../mock-data/category';
import { CATEGORIES } from '../mock-data/mock-categories';
import { EventEmitter } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {

  constructor() { }

  getCategories(): Category[] {
    return CATEGORIES;
  }
}

