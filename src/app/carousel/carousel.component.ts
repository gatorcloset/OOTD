import { Component, Input, OnInit } from '@angular/core';
import { CarouselService } from '../services/carousel.service';
import { Item } from '../mock-data/item';
import { Outfit } from '../mock-data/outfit';
import { Router } from '@angular/router';
import { UserService } from '../services/user.service';
import { User } from '../mock-data/user';

interface carouselImage {
  imageSrc: string;
  imageAlt: string;
}

@Component({
  selector: 'app-carousel',
  templateUrl: './carousel.component.html',
  styleUrls: ['./carousel.component.css']
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

  constructor(private carouselService: CarouselService, private router: Router, private userService: UserService) { }

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

  createOutfit(name: string, top: number, bottom: number, shoes: number, accessory: number) {
    const outfit: Outfit = {
      user_id: this.userService.authUser?.ID, 
      Name: name,
      Tops: this.tops[top],
      Bottoms: this.bottoms[bottom],
      Shoes: this.shoes[shoes],
      Accessories: this.accessories[accessory]
    }

    this.carouselService.createOutfit(outfit).subscribe(
      res => {
        console.log(res)
        this.router.navigateByUrl('/outfits');
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
