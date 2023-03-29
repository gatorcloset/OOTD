import { Component, Input, OnInit } from '@angular/core';

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

  @Input() tops: carouselImage[] = []
  @Input() bottoms: carouselImage[] = []
  @Input() shoes: carouselImage[] = []
  @Input() indicators = true;
  @Input() controls = true;

  selectedIndex = 0;

  ngOnInit():void {
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
