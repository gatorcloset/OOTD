import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Item } from '../mock-data/item';

@Injectable({
  providedIn: 'root'
})
export class NewItemService {
  url: string = 'http://localhost:9000/item';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'multipart/form-data' })
  }

  formData: FormData = new FormData();

  constructor(private http: HttpClient) { }

  set(key: string, value: any) {
    this.formData.set(key, value);
    // console.log(value);
  }

  createItem(): Observable<Item> {
    console.log(this.formData.get('name'));
    console.log(this.formData.get('category'));
    console.log(this.formData.get('image'));
    console.log(this.formData.get('id'));
    return this.http.post<Item>(this.url, this.formData);
  }

}
