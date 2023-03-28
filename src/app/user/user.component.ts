import { Component } from '@angular/core';
import { User } from '../mock-data/user';
import { UserService } from '../services/user.service';


@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent {
  user: User = <User>{};

  constructor(private userService: UserService) {}

  ngOnInit(): void {
    this.userService.getUser().subscribe(
      user => this.user = user,
      error => console.error(error)
    )
  }
}
