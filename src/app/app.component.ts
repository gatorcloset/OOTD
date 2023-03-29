import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'ootd-app';
  constructor(private router: Router) {}

  noNavbar(): boolean {
    return this.router.url.includes('login') || this.router.url.includes('signup') || this.router.url === '/';
  }

  inBuilder(): boolean {
    return this.router.url.includes('builder');
  }

  //sample images for outfit builder
  tops = [
    {
      imageSrc: 'assets/builder-sample-pics/tops/black-blouse.webp',
      imageAlt: 'black-blouse',
    },
    {
      imageSrc: 'assets/builder-sample-pics/tops/green-cami.avif',
      imageAlt: 'green-cami',
    },
    {
      imageSrc: 'assets/builder-sample-pics/tops/pink-sweatshirt.avif',
      imageAlt: 'pink-sweatshirt',
    },
    {
      imageSrc: 'assets/builder-sample-pics/tops/purple-tank.avif',
      imageAlt: 'purple-tank',
    },
    {
      imageSrc: 'assets/builder-sample-pics/tops/white-sweatshirt.avif',
      imageAlt: 'white-sweatshirt',
    }
  ]

  bottoms = [
    {
      imageSrc: 'assets/builder-sample-pics/bottoms/biker-shorts.avif',
      imageAlt: 'biker-shorts',
    },
    {
      imageSrc: 'assets/builder-sample-pics/bottoms/black-pants.avif',
      imageAlt: 'black-pants',
    },
    {
      imageSrc: 'assets/builder-sample-pics/bottoms/black-shorts.avif',
      imageAlt: 'black-shorts',
    },
    {
      imageSrc: 'assets/builder-sample-pics/bottoms/mom-jeans.avif',
      imageAlt: 'mom-jeans',
    },
    {
      imageSrc: 'assets/builder-sample-pics/bottoms/pink-skirt.avif',
      imageAlt: 'pink-skirt',
    }
  ]

  shoes = [
    {
      imageSrc: 'assets/builder-sample-pics/footwear/shoes1.webp',
      imageAlt: 'biker-shorts',
    },
    {
      imageSrc: 'assets/builder-sample-pics/footwear/shoes2.webp',
      imageAlt: 'black-pants',
    },
    {
      imageSrc: 'assets/builder-sample-pics/footwear/shoes3.jpg',
      imageAlt: 'black-shorts',
    },
    {
      imageSrc: 'assets/builder-sample-pics/footwear/shoes-4.jpg',
      imageAlt: 'mom-jeans',
    },
    {
      imageSrc: 'assets/builder-sample-pics/footwear/shoes5.jpg',
      imageAlt: 'pink-skirt',
    }
  ]
}
