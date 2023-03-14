import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class NewItemService {
  formData: FormData = new FormData();

  constructor(private http: HttpClient) { }

  set(key: string, value: any) {
    this.formData.set(key, value);
    console.log(this.formData);
  }

}
