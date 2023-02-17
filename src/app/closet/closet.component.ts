import { Component, Output, EventEmitter } from '@angular/core';
import { Category } from '../mock-data/category';
import { CategoryService } from '../services/category.service';

@Component({
  selector: 'app-closet',
  templateUrl: './closet.component.html',
  styleUrls: ['./closet.component.css']
})
export class ClosetComponent {
  categories: Category[] = [];

  constructor(private categoryService: CategoryService) {}

  getCategories(): void {
    this.categories = this.categoryService.getCategories();
  }

  ngOnInit(): void {
    this.getCategories();
  }
}
