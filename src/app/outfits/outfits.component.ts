import { Component } from '@angular/core';
import { CarouselService } from '../services/carousel.service';
import { Outfit } from '../mock-data/outfit';
import { Router } from '@angular/router';

@Component({
  selector: 'app-outfits',
  templateUrl: './outfits.component.html',
  styleUrls: ['./outfits.component.css']
})
export class OutfitsComponent {
  allOutfits: Outfit[] = [];

  constructor(private carouselService: CarouselService, private router: Router) { }

  getOutfits() {
    this.carouselService.getOutfits().subscribe(
      res => {
        this.allOutfits = res;
        console.log(res)
      },
      err => console.log(err)
    )
  }

  ngOnInit() {
    this.getOutfits();
  }

}
