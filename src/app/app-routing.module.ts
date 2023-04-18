import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EntranceComponent } from './entrance/entrance.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { ClosetComponent } from './closet/closet.component';
import { UserComponent } from './user/user.component';
import { HomeComponent } from './home/home.component';
import { NewItemComponent } from './new-item/new-item.component';
import { CarouselComponent } from './carousel/carousel.component';
import { AuthGuard } from './services/auth-guard.service';
import { ItemsComponent } from './items/items.component';
import { OutfitsComponent } from './outfits/outfits.component';


const routes: Routes = [
  { path: 'closet', component: ClosetComponent, canActivate: [AuthGuard] },
  { path: 'closet/:name', component: ItemsComponent, canActivate: [AuthGuard] },
  { path: 'outfits', component: OutfitsComponent, canActivate: [AuthGuard]},
  { path: '', component: EntranceComponent },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
  { path: 'user', component: UserComponent, canActivate: [AuthGuard] },
  { path: 'home', component: HomeComponent},
  { path: 'add', component: NewItemComponent, canActivate: [AuthGuard] },
  { path: 'builder', component: CarouselComponent, canActivate: [AuthGuard] }
];

/*
const routes: Routes = [
  { path: 'closet', component: ClosetComponent},
  { path: 'closet/:name', component: ItemsComponent},
  { path: 'outfits', component: OutfitsComponent},
  { path: '', component: EntranceComponent },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
  { path: 'user', component: UserComponent},
  { path: 'home', component: HomeComponent},
  { path: 'add', component: NewItemComponent},
  { path: 'builder', component: CarouselComponent}
];
*/

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
