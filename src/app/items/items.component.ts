import { Component, Inject } from '@angular/core';
import { Item } from 'src/app/mock-data/item';
import { ItemService } from 'src/app/services/item.service';
import { ActivatedRoute } from '@angular/router';
import {MatDialog, MAT_DIALOG_DATA, MatDialogRef} from '@angular/material/dialog';
import { NewItemService } from '../services/new-item.service';
import { UserService } from '../services/user.service';
import { Router } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ChangeDetectorRef } from '@angular/core';

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
  constructor(private itemService: ItemService, private activatedRoute: ActivatedRoute, public dialog: MatDialog) {}

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

  openDialog(item: Item) {
    const dialogRef = this.dialog.open(ItemsDialogComponent, {
      data: {name: item.name, category: item.category, image: item.image, ID: item.ID},
    });
  }

  async ngOnInit() {
    // Retrieves array of all mock items
    const res = await this.getItems();
  
    // Retrieves the name element of the router
    this.selectedCategory = this.activatedRoute.snapshot.paramMap.get('name')!;
    // console.log(this.selectedCategory);
    // Sets the array of selected items = to the original items array, but filtered
    this.selectedItems = this.items.filter(x => x.category.toLowerCase() === this.selectedCategory);
    // console.log(this.selectedItems);
  }
}

@Component({
  selector: 'app-items-dialog',
  templateUrl: 'items.dialog.component.html',
  styleUrls: ['./items.dialog.component.css']
})
export class ItemsDialogComponent {
  item: Item;
  selectedFile: File | null = null;
  imageUrl: string = "";

  constructor(
    public dialogRef: MatDialogRef<ItemsDialogComponent>,
    private newItemService: NewItemService,
    private userService: UserService,
    private router: Router,
    private route: ActivatedRoute,
    private cdr: ChangeDetectorRef,
    @Inject(MAT_DIALOG_DATA) public data: Item,
  ) {
    this.item = data; // Assign 'item' property from 'data' object
    this.imageUrl = this.item.image;
  }

  onChangeFile(event: any) {
    if (event.target.files.length > 0) {
      this.selectedFile = event.target.files[0];

      // call the new-item service function that sets the image field to selected file
      this.newItemService.set('image', this.selectedFile);

      // sets image to be displayed
      if (this.selectedFile) {
        const reader = new FileReader();
        reader.readAsDataURL(this.selectedFile);
        reader.onload = () => {
          this.imageUrl = reader.result as string;
        };
      }
    }
    
  }

  onSave(itemID: number, name: string, category: string, event: Event) {
    category = category || this.item.category;

    // Populate FormData object based on input values
    this.newItemService.set('name', name);
    this.newItemService.set('category', category);
    this.newItemService.set('id', itemID);

    // Call PUT function, using the itemID as a parameter
    this.newItemService.updateItem(itemID).subscribe(
      res => {
        console.log(res)
        this.onNoClick();
        this.router.navigate(['/closet', category]).then(() => {
          this.cdr.detectChanges();
        })
        // this.router.navigateByUrl(`/closet/${category}`) => does not work :(
        this.router.navigateByUrl('/closet')
      },
      err => console.log(err)
    )
  }

  onNoClick(): void {
    this.dialogRef.close();
  }
}