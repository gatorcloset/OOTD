import { Component } from '@angular/core';
import { User } from '../mock-data/data';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent {
  newUser: User = <User>{};

  constructor(private userService: UserService) { }

  createUser(firstname: string, lastname: string, username: string, password: string) {
    // Trim any added spaces after the name
    firstname = firstname.trim();
    lastname = lastname.trim();
    username = username.trim();

    this.userService.createUser({ firstname, lastname, username, password } as User).subscribe(
      newUser => this.newUser = newUser
    )

  }
}
