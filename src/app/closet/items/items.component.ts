import { Component } from '@angular/core';
import { ITEMS } from 'src/app/mock-items';

@Component({
  selector: 'app-items',
  templateUrl: './items.component.html',
  styleUrls: ['./items.component.css']
})
export class ItemsComponent {
  items = ITEMS;
}
