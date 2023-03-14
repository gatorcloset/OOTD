import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NavbarComponent } from './navbar/navbar.component';
import { ClosetComponent } from './closet/closet.component';
import { ItemsComponent } from './items/items.component';

import {MatToolbarModule} from '@angular/material/toolbar';
import {MatCardModule} from '@angular/material/card';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatDividerModule} from '@angular/material/divider';
import {MatGridListModule} from '@angular/material/grid-list';
import { EntranceComponent } from './entrance/entrance.component';
import { LoginComponent } from './login/login.component';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatInputModule} from '@angular/material/input';
import { SignupComponent } from './signup/signup.component';

import { ClosetModule } from './closet/closet.module';
import { UserComponent } from './user/user.component';
import { HomeComponent } from './home/home.component';
import { NewItemComponent } from './new-item/new-item.component';
import { MatChipsModule } from '@angular/material/chips';
import { ReactiveFormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { UploadImageComponent } from './upload-image/upload-image.component';
import {MatSelectModule} from '@angular/material/select';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    ClosetComponent,
    ItemsComponent,
    EntranceComponent,
    LoginComponent,
    SignupComponent,
    UserComponent,
    HomeComponent,
    NewItemComponent,
    UploadImageComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatToolbarModule,
    MatCardModule,
    MatIconModule,
    MatButtonModule,
    MatDividerModule,
    MatGridListModule,
    //RouterModule.forRoot(appRoute),
    MatFormFieldModule,
    MatInputModule,
    ClosetModule,
    AppRoutingModule,
    HttpClientModule,
    MatChipsModule,
    ReactiveFormsModule,
    MatAutocompleteModule,
    MatSelectModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
