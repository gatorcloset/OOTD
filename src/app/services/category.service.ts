import { Injectable } from '@angular/core';
import { Category } from '../mock-data/category';
import { CATEGORIES } from '../mock-data/default-categories';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {

  constructor() { }

  getCategories(): Category[] {
    return CATEGORIES;
  }
}

