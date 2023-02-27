import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../mock-data/data';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiURL = 'http://localhost:9000';

  constructor(private http: HttpClient) { }

  getUser(): Observable<User> {
    const url = `${this.apiURL}/users/1`;
    return this.http.get<User>(url);
  }
}
