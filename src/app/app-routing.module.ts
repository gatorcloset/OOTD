import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EntranceComponent } from './entrance/entrance.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { ClosetComponent } from './closet/closet.component';
import { UserComponent } from './user/user.component';
import { CarouselComponent } from './carousel/carousel.component';

const routes: Routes = [
  { path: 'closet', component: ClosetComponent },
  { path: '', component: EntranceComponent },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
  { path: 'user', component: UserComponent},
  { path: 'builder', component: CarouselComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
