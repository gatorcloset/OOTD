import { Injectable } from '@angular/core';
import { Item } from '../mock-data/item';
import { ITEMS } from '../mock-data/mock-items';

@Injectable({
  providedIn: 'root'
})
export class ItemService {

  constructor() { }

  getItems(): Item[] {
    return ITEMS;
  }
}
