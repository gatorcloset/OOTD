import { Component, Input } from '@angular/core';
import { Item } from 'src/app/mock-data/item';
import { ItemService } from 'src/app/services/item.service';

@Component({
  selector: 'app-items',
  templateUrl: './items.component.html',
  styleUrls: ['./items.component.css']
})
export class ItemsComponent {
  // @Input() selectedCategory?: Category;

  items: Item[] = [];

  constructor(private itemService: ItemService) {}

  getItems(): void {
    this.items = this.itemService.getItems();
  }

  ngOnInit(): void {
    this.getItems();
  }
}
