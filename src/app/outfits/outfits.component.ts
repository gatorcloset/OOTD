import { Component, Inject, Input, OnInit } from '@angular/core';
import { CarouselService } from '../services/carousel.service';
import { Outfit } from '../mock-data/outfit';
import { Router } from '@angular/router';
import { MatDialog, MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { Item } from '../mock-data/item';

@Component({
  selector: 'app-outfits',
  templateUrl: './outfits.component.html',
  styleUrls: ['./outfits.component.css']
})
export class OutfitsComponent {
  allOutfits: Outfit[] = [];

  constructor(private carouselService: CarouselService, private router: Router, public dialog: MatDialog) { }

  getOutfits() {
    this.carouselService.getOutfits().subscribe(
      res => {
        this.allOutfits = res;
        console.log(res)
      },
      err => console.log(err)
    )
  }

  openDialog(outfit: Outfit) {
    const dialogRef = this.dialog.open(OutfitsDialogComponent, {
      data: {ID: outfit.ID, Name: outfit.Name, Bottoms: outfit.Bottoms, Tops: outfit.Tops, Shoes: outfit.Shoes, Accessories: outfit.Accessories},
    });
  }

  ngOnInit() {
    this.getOutfits();
  }

}

@Component({
  selector: 'app-outfits-dialog',
  templateUrl: './outfits.dialog.component.html',
  styleUrls: ['./outfits.dialog.component.css']
})
export class OutfitsDialogComponent {

  @Input() indicators = true;
  @Input() controls = true;

  tops: Item[] = [];
  bottoms: Item[] = [];
  onePieces: Item[] = [];
  shoes: Item[] = [];
  accessories: Item[] = [];


  selectedIndex = 0;
  topsIndex = 0;
  bottomsIndex = 0;
  shoesIndex = 0;
  accessoriesIndex = 0;

  constructor(
    public dialogRef: MatDialogRef<OutfitsDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: Outfit,
    private carouselService: CarouselService,
    private router: Router
  ) { }

  onNoClick(): void {
    this.dialogRef.close();
  }

  getItemByCategory() {
    this.carouselService.getItemByCategory('tops').subscribe(
      res => {
        this.tops = res;
      },
      err => {
        console.log(err);
      }
    )
  
    this.carouselService.getItemByCategory('bottoms').subscribe(
      res => {
        this.bottoms = res;
      },
      err => {
        console.log(err);
      }
    )

    this.carouselService.getItemByCategory('one-pieces').subscribe(
      res => {
        this.onePieces = res;
      },
      err => {
        console.log(err);
      }
    )

    this.carouselService.getItemByCategory('shoes').subscribe(
      res => {
        this.shoes = res;
      },
      err => {
        console.log(err);
      }
    )

    this.carouselService.getItemByCategory('accessories').subscribe(
      res => {
        this.accessories = res;
      },
      err => {
        console.log(err);
      }
    )
  
  }

  getOutfits() {
    this.carouselService.getOutfits().subscribe(
      res => console.log(res),
      err => console.log(err)
    )
  }

  updateOutfit(ID: number, name: string, top: number, bottom: number, shoes: number, accessory: number) {
    const outfit: Outfit = {
      ID: ID,
      Name: name,
      Tops: this.tops[top],
      TopID: this.tops[top].ID,
      Bottoms: this.bottoms[bottom],
      BottomID: this.bottoms[bottom].ID,
      Shoes: this.shoes[shoes],
      ShoesID: this.shoes[shoes].ID,
      Accessories: this.accessories[accessory],
      AccessoriesID: this.accessories[accessory].ID
    }

    console.log(outfit);

    this.carouselService.updateOutfit(outfit).subscribe(
      res => {
        console.log(res);
        this.onNoClick();
        this.router.navigateByUrl('/closet');
      },
      err => console.log(err)
    )
  }

  deleteOutfit(outfit: Outfit) {
    this.carouselService.deleteOutfit(outfit).subscribe(
      res => {
        console.log(res);
        this.onNoClick();
        this.router.navigateByUrl('/closet');
      },
      err => console.log(err)
    )
  }

  ngOnInit():void {
    this.getItemByCategory();
  }

  // sets index of image on dot click
  selectImage(index: number): void {
    this.selectedIndex = index;
  }

  topsPrevClick(): void {
    if(this.topsIndex === 0) {
      this.topsIndex = this.tops.length - 1;
    } else {
      this.topsIndex--;
    }
  }

  topsNextClick(): void {
    if(this.topsIndex === this.tops.length-1) {
      this.topsIndex = 0;
    } else {
      this.topsIndex++;
    }
  }

  bottomsPrevClick(): void {
    if(this.bottomsIndex === 0) {
      this.bottomsIndex = this.bottoms.length - 1;
    } else {
      this.bottomsIndex--;
    }
  }

  bottomsNextClick(): void {
    if(this.bottomsIndex === this.bottoms.length-1) {
      this.bottomsIndex = 0;
    } else {
      this.bottomsIndex++;
    }
  }

  shoesPrevClick(): void {
    if(this.shoesIndex === 0) {
      this.shoesIndex = this.shoes.length - 1;
    } else {
      this.shoesIndex--;
    }
  }

  shoesNextClick(): void {
    if(this.shoesIndex === this.shoes.length-1) {
      this.shoesIndex = 0;
    } else {
      this.shoesIndex++;
    }
  }

  accessoriesPrevClick(): void {
    if(this.accessoriesIndex === 0) {
      this.accessoriesIndex = this.accessories.length - 1;
    } else {
      this.accessoriesIndex--;
    }
    console.log(this.accessoriesIndex);
  }

  accessoriesNextClick(): void {
    if(this.accessoriesIndex === this.accessories.length-1) {
      this.accessoriesIndex = 0;
    } else {
      this.accessoriesIndex++;
    }
    console.log(this.accessoriesIndex);
  }

}


