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
import { FormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { UploadImageComponent } from './upload-image/upload-image.component';
import { MatSelectModule } from '@angular/material/select';
// import { MaterialFileInputModule } from 'ngx-material-file-input';
import { CarouselModule } from './carousel/carousel.module';
import { CarouselComponent } from './carousel/carousel.component';
import { AuthGuard } from './services/auth-guard.service';
import { UserService } from './services/user.service';
import {MatDialogModule} from '@angular/material/dialog';
import { ItemsDialogComponent } from './items/items.component';

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
    UploadImageComponent,
    ItemsDialogComponent
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
    FormsModule,
    MatAutocompleteModule,
    MatSelectModule,
    // MaterialFileInputModule,
    CarouselModule,
    MatDialogModule
  ],
  providers: [UserService, AuthGuard],
  bootstrap: [AppComponent]
})
export class AppModule { }
