import { Component } from '@angular/core';
import { NewItemService } from '../services/new-item.service';

@Component({
  selector: 'app-upload-image',
  templateUrl: './upload-image.component.html',
  styleUrls: ['./upload-image.component.css']
})
export class UploadImageComponent {
  selectedFile: File | null = null;
  imageUrl: string = "";

  constructor(private newItemService: NewItemService) {}

  onChangeFile(event: any) {
    if (event.target.files.length > 0) {
      this.selectedFile = event.target.files[0];

      // call the new-item service function that sets the image field to selected file
      this.newItemService.set('image', this.selectedFile);

      // sets image to be displayed
      if (this.selectedFile) {
        const reader = new FileReader();
        reader.readAsDataURL(this.selectedFile);
        reader.onload = () => {
          this.imageUrl = reader.result as string;
        };
      }
    }
    
  }

}
