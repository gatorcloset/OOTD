import { Injectable } from '@angular/core';
import { Item } from '../mock-data/item';
import { UserService } from './user.service';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ItemService {
  private apiURL = 'http://localhost:9000';

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  }

  constructor(private userService: UserService, private http: HttpClient) { }

  getItems(): Observable<Item[]> {
    const url = `${this.apiURL}/users/${this.userService.authUser?.ID}/items`;
    return this.http.get<Item[]>(url, this.httpOptions);
  }
}
