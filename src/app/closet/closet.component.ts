import { Component, Output, EventEmitter } from '@angular/core';
import { Category } from '../mock-data/category';
import { CategoryService } from '../services/category.service';

export interface User {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  firstname: string;
  lastname: string;
  email: string;
  password: string;
}

@Component({
  selector: 'app-closet',
  templateUrl: './closet.component.html',
  styleUrls: ['./closet.component.css']
})
export class ClosetComponent {
  categories: Category[] = [];
  testUsers: User[] = [];

  constructor(private categoryService: CategoryService) {}

  getCategories(): void {
    this.categories = this.categoryService.getCategories();
  }

  ngOnInit(): void {
    this.getCategories();
  
    this.categoryService.getAPI().subscribe(
      users => this.testUsers = users,
      error => console.error(error)
    )
  
  }
}
