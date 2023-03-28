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
  images = [
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
}
