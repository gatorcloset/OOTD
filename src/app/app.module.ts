import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NavbarComponent } from './navbar/navbar.component';
import { ClosetComponent } from './closet/closet.component';

import {MatToolbarModule} from '@angular/material/toolbar';
import {MatCardModule} from '@angular/material/card';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatDividerModule} from '@angular/material/divider';
import {MatGridListModule} from '@angular/material/grid-list';
import { ItemsComponent } from './closet/items/items.component';
import { EntranceComponent } from './entrance/entrance.component';
import { LoginComponent } from './login/login.component';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import { SignupComponent } from './signup/signup.component';

const appRoute: Routes = [
  {path: 'closet', component: ClosetComponent},
  {path: 'closet', children: [
    {path: 'jeans', component: ItemsComponent} // TO DO: MAKE PATH CUSTOM TO WHATEVER IS CLICKED
  ]}
]

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    ClosetComponent,
    ItemsComponent,
    EntranceComponent,
    LoginComponent,
    SignupComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    MatCardModule,
    MatIconModule,
    MatButtonModule,
    MatDividerModule,
    MatGridListModule,
    RouterModule.forRoot(appRoute),
    MatFormFieldModule,
    MatInputModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
