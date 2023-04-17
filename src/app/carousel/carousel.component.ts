import { Component, Input, OnInit } from '@angular/core';
import { CarouselService } from '../services/carousel.service';
import { Item } from '../mock-data/item';

interface carouselImage {
  imageSrc: string;
  imageAlt: string;
}

@Component({
  selector: 'app-carousel',
  templateUrl: './carousel.component.html',
  styleUrls: ['./carousel.component.scss']
})

export class CarouselComponent implements OnInit{

  /*
  @Input() tops: carouselImage[] = []
  @Input() bottoms: carouselImage[] = []
  @Input() shoes: carouselImage[] = []
  */

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

  constructor(private carouselService: CarouselService) { }

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

  saveOutfit(top: number, bottom: number, shoes: number, accessory: number) {
    // Create an array to store the items in the outfit
    const outfit: Item[] = [];

    // Retrieve items from respective arrays based on indices
    if (top >= 0 && top < this.tops.length) {
      outfit.push(this.tops[top]);
    }

    if (bottom >= 0 && bottom < this.bottoms.length) {
      outfit.push(this.bottoms[bottom]);
    }

    /*
    if (onePiece >= 0 && onePiece < this.onePieces.length) {
      outfit.push(this.onePieces[onePiece]);
    }
    */

    if (accessory >= 0 && accessory < this.accessories.length) {
      outfit.push(this.accessories[accessory]);
    }

    if (shoes >= 0 && shoes < this.shoes.length) {
      outfit.push(this.shoes[shoes]);
    }

    console.log(outfit);

    this.carouselService.saveOutfit(outfit).subscribe(
      res => {
        console.log("this is the result");
        console.log(res)
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
  }

  accessoriesNextClick(): void {
    if(this.accessoriesIndex === this.accessories.length-1) {
      this.accessoriesIndex = 0;
    } else {
      this.accessoriesIndex++;
    }
  }
}
