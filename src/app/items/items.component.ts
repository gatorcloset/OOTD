import { Component, Input } from '@angular/core';
import { Item } from 'src/app/mock-data/item';
import { ItemService } from 'src/app/services/item.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-items',
  templateUrl: './items.component.html',
  styleUrls: ['./items.component.css']
})
export class ItemsComponent {
  // @Input() selectedCategory?: Category;

  items: Item[] = [];
  selectedCategory: string = "";
  selectedItems: Item[] = [];

  // Creates an instance of the ItemService and CategoryService
  constructor(private itemService: ItemService, private activatedRoute: ActivatedRoute) {}

  getItems(): void {
    // Populates the items array
    this.items = this.itemService.getItems(); 
  }

  ngOnInit(): void {
    // Retrieves array of all mock items
    this.getItems();
  
    // Retrieves the name element of the router
    this.selectedCategory = this.activatedRoute.snapshot.paramMap.get('name')!;
    // Sets the array of selected items = to the original items array, but filtered
    this.selectedItems = this.itemService.getItems().filter(x => x.category.toLowerCase() === this.selectedCategory);
  }
}
