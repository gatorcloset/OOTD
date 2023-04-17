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

  saveOutfit(outfit: Item[]): Observable<Outfit> {
    const url = `${this.apiURL}/outfit`;
    return this.http.post<Outfit>(url, outfit);
  }

  getOutfits(): Observable<Outfit[]> {
    const url = `${this.apiURL}/outfit`;
    return this.http.get<Outfit[]>(url);
  }


}
