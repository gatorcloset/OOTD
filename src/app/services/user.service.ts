import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../mock-data/data';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiURL = 'http://localhost:9000/users';

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  }

  constructor(private http: HttpClient) { }

  getUser(): Observable<User> {
    const url = `${this.apiURL}/1`;
    return this.http.get<User>(url);
  }

  createUser(user: User): Observable<User> {
    const url =  `${this.apiURL}`;
    return this.http.post<User>(url, user, this.httpOptions);
  }
}
