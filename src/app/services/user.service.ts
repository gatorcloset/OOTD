import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import { tap, catchError } from 'rxjs';
import { User } from '../mock-data/user';
import { LoginRequest } from '../mock-data/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiURL = 'http://localhost:9000/users';
  
  // Authenticated user information
  authenticated = false;
  authUser?: User = <User>{};

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  }
  
  constructor(private http: HttpClient) { }

  getUser(): Observable<User> {
    const url = `${this.apiURL}/${this.authUser?.ID}`;
    return this.http.get<User>(url);
  }

  createUser(user: User): Observable<User> {
    const url =  `${this.apiURL}`;
    return this.http.post<User>(url, user, this.httpOptions);
  }

  loginUser(user: LoginRequest): Observable<User> {
    const url = "http://localhost:9000/login";
    const options = { ...this.httpOptions, withCredentials: true };
    return this.http.post<User>(url, user, options)
  }
}
