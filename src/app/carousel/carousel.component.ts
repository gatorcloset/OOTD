import { Component, Input, OnInit } from '@angular/core';
import { CarouselService } from '../services/carousel.service';
import { Item } from '../mock-data/item';
import { Outfit } from '../mock-data/outfit';

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

  getOutfits() {
    this.carouselService.getOutfits().subscribe(
      res => console.log(res),
      err => console.log(err)
    )
  }

  saveOutfit(name: string, top: number, bottom: number, shoes: number) {
    const outfit: Outfit = {
      Name: name,
      Tops: this.tops[top],
      Bottoms: this.bottoms[top],
      Shoes: this.shoes[top]
    }

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
