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

  getItems() {
    return new Promise(resolve => {
      this.itemService.getItems().subscribe(
        res => {
          this.items = res;
          console.log(res);
          resolve(res);
        },
        err => {
          console.log(err);
          resolve(err);
        }
      )
    })
    
  }

  async ngOnInit() {
    // Retrieves array of all mock items
    const res = await this.getItems();
  
    // Retrieves the name element of the router
    this.selectedCategory = this.activatedRoute.snapshot.paramMap.get('name')!;
    console.log(this.selectedCategory);
    // Sets the array of selected items = to the original items array, but filtered
    this.selectedItems = this.items.filter(x => x.category.toLowerCase() === this.selectedCategory);
    console.log(this.selectedItems);
  }
}
