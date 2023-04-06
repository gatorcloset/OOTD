import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../mock-data/user';
import { LoginRequest } from '../mock-data/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiURL = 'http://localhost:9000/users';
  authUser: User = <User>{};

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  }
  
  /*
  session = {
    headers: new HttpHeaders({ 
      'Content-Type':  'application/json'
    }),
    withCredentials: true
  };
  */

  constructor(private http: HttpClient) { }

  setAuthUser(user: User) {
    this.authUser = user;
  }

  getAuthUser(): User {
    return this.authUser;
  }

  getUser(): Observable<User> {
    const userID = sessionStorage.getItem('userID'); // THIS IS CURRENTLY NULL
    const url = `${this.apiURL}/${userID}`;
    return this.http.get<User>(url);
  }

  createUser(user: User): Observable<User> {
    const url =  `${this.apiURL}`;
    return this.http.post<User>(url, user, this.httpOptions);
  }

  loginUser(user: LoginRequest): Observable<User> {
    const url = "http://localhost:9000/login";
    const options = { ...this.httpOptions, withCredentials: true };
    return this.http.post<User>(url, user, options);
  }
}
