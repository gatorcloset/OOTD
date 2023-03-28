import {COMMA, ENTER} from '@angular/cdk/keycodes';
import {Component, ElementRef, ViewChild} from '@angular/core';
import {FormControl} from '@angular/forms';
import {MatAutocompleteSelectedEvent} from '@angular/material/autocomplete';
import {MatChipInputEvent} from '@angular/material/chips';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import { Tag } from '../mock-data/tag';
import { NewItemService } from '../services/new-item.service';
import { TagService } from '../services/tag.service';


@Component({
  selector: 'app-new-item',
  templateUrl: './new-item.component.html',
  styleUrls: ['./new-item.component.css']
})
export class NewItemComponent {
  separatorKeysCodes: number[] = [ENTER, COMMA];
  tagCtrl = new FormControl('');
  filteredTags: Observable<string[]>;
  allTags: Tag[] = []; // All tags from database
  tags: string[] = []; // Tags selected by the user in the chips component
  selectedCat: string = "";

  @ViewChild('tagInput', {static: true}) tagInput!: ElementRef<HTMLInputElement>;

  constructor(private tagService: TagService, private newItem: NewItemService) {
    this.filteredTags = this.tagCtrl.valueChanges.pipe(
      startWith(null),
      map((tag: string | null) => (tag ? this._filter(tag).map(tag => tag.name) : this.allTags.map(tag => tag.name)),
    ));
  }

  add(event: MatChipInputEvent): void {
    const value = (event.value || '').trim();

    // Add our fruit
    if (value) {
      this.tags.push(value);
    }

    // Clear the input value
    event.chipInput!.clear();

    this.tagCtrl.setValue(null);
  }

  remove(tag: string): void {
    const index = this.tags.indexOf(tag);

    if (index >= 0) {
      this.tags.splice(index, 1);
    }
  }

  selected(event: MatAutocompleteSelectedEvent): void {
    this.tags.push(event.option.viewValue);
    this.tagInput.nativeElement.value = '';
    this.tagCtrl.setValue(null);
  }

  private _filter(value: string): Tag[] {
    const filterValue = value.toLowerCase();

    return this.allTags.filter(tag => tag.name.toLowerCase().includes(filterValue));
  }

  getTags(): void {
    this.allTags = this.tagService.getTags();
  }

  onSelected(value: string) {
    this.selectedCat = value;
  }

  onSubmit(name: string, category: string, event: Event) {
    event.preventDefault();
    // Populate FormData object based on input values
    this.newItem.set('name', name);
    this.newItem.set('category', category);

    // Call POST function
    this.newItem.createItem().subscribe(
      res => console.log(res),
      error => console.error(error)
    )
  }

  ngOnInit(): void {
    this.getTags();
  }

}
