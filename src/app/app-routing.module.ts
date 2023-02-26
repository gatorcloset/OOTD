import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EntranceComponent } from './entrance/entrance.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { ClosetComponent } from './closet/closet.component';
// import { ItemsComponent } from './closet/items/items.component';

const routes: Routes = [
  {path: 'closet', component: ClosetComponent},
  /*
  {path: 'closet', children: [
    {path: 'jeans', component: ItemsComponent} // TO DO: MAKE PATH CUSTOM TO WHATEVER IS CLICKED
  ]}
  */

  {path: '', component: EntranceComponent},
  {path: 'login', component: LoginComponent},
  {path: 'signup', component: SignupComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
