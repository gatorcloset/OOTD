import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EntranceComponent } from './entrance/entrance.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { ClosetComponent } from './closet/closet.component';
import { UserComponent } from './user/user.component';
import { HomeComponent } from './home/home.component';
import { NewItemComponent } from './new-item/new-item.component';

const routes: Routes = [
  { path: 'closet', component: ClosetComponent },
  { path: '', component: EntranceComponent },
  { path: 'login', component: LoginComponent },
  { path: 'signup', component: SignupComponent },
  { path: 'user', component: UserComponent},
  { path: 'home', component: HomeComponent},
  { path: 'add-item', component: NewItemComponent}

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
