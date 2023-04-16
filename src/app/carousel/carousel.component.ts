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

  saveOutfit(top: number, bottom: number, shoes: number) {
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

    if (accessory >= 0 && accessory < this.accessories.length) {
      outfit.push(this.accessories[accessory]);
    }
    */

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
    if(this.selectedIndex === 0) {
      this.selectedIndex = this.tops.length - 1;
    } else {
      this.selectedIndex--;
    }
  }

  topsNextClick(): void {
    if(this.selectedIndex === this.tops.length-1) {
      this.selectedIndex = 0;
    } else {
      this.selectedIndex++;
    }
  }

  bottomsPrevClick(): void {
    if(this.selectedIndex === 0) {
      this.selectedIndex = this.bottoms.length - 1;
    } else {
      this.selectedIndex--;
    }
  }

  bottomsNextClick(): void {
    if(this.selectedIndex === this.bottoms.length-1) {
      this.selectedIndex = 0;
    } else {
      this.selectedIndex++;
    }
  }

  shoesPrevClick(): void {
    if(this.selectedIndex === 0) {
      this.selectedIndex = this.shoes.length - 1;
    } else {
      this.selectedIndex--;
    }
  }

  shoesNextClick(): void {
    if(this.selectedIndex === this.shoes.length-1) {
      this.selectedIndex = 0;
    } else {
      this.selectedIndex++;
    }
  }
}
