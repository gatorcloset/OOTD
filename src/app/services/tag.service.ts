import { Injectable } from '@angular/core';
import { Tag } from '../mock-data/tag';
import { TAGS } from '../mock-data/tag';

@Injectable({
  providedIn: 'root'
})
export class TagService {

  constructor() { }

  getTags(): Tag[] {
    return TAGS;
  }
  
}
