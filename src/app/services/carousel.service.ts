import { Injectable } from '@angular/core';
import { Item } from '../mock-data/item';
import { Outfit } from '../mock-data/outfit';
import { UserService } from './user.service';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CarouselService {
  private apiURL = 'http://localhost:9000';

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  }

  constructor(private userService: UserService, private http: HttpClient) { }

  getItemByCategory(category: string): Observable<Item[]> {
    const url = `${this.apiURL}/users/${this.userService.authUser?.ID}/category/${category}`;
    return this.http.get<Item[]>(url, this.httpOptions);
  }

  createOutfit(outfit: Outfit): Observable<Outfit> {
    const url = `${this.apiURL}/outfit`;
    return this.http.post<Outfit>(url, outfit);
  }

  getOutfits(): Observable<Outfit[]> {
    const url = `${this.apiURL}/users/${this.userService.authUser?.ID}/outfits`;
    return this.http.get<Outfit[]>(url);
  }

  updateOutfit(outfit: Outfit): Observable<Outfit> {
    console.log(outfit);
    const url = `${this.apiURL}/outfit/${outfit.ID}`;
    return this.http.put<Outfit>(url, outfit);
  }

  deleteOutfit(outfit: Outfit): Observable<Outfit> {
    const url = `${this.apiURL}/outfit/${outfit.ID}`;
    return this.http.delete<Outfit>(url);
  }


}
