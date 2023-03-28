import { Component } from '@angular/core';
import {FormControl, Validators} from '@angular/forms';
import { LoginRequest } from '../mock-data/user';
import { UserService } from '../services/user.service';
import { User } from '../mock-data/user';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent {
  formValid = false;
  username = new FormControl('', [Validators.required, Validators.pattern('\\S+')]);
  password = new FormControl('', [Validators.required]);
  loginError: string = "";

  constructor(private userService: UserService) {

  }

  userErrorMessage() {
    if (this.username.hasError('required')) {
      return 'You must enter a username';
    }

    if (this.username.hasError('pattern')) {
      return 'Not a valid username'
    }
    
    return '';
  }

  passErrorMessage() {
    if (this.password.hasError('required')) {
      return 'You must enter a password'
    }
    
    return '';
  }

  login(username: string, password: string) {
    this.userService.loginUser({ username, password } as LoginRequest).subscribe(
      res => {
        // Successful login => save authenticated user
        this.userService.setAuthUser(res);
        this.loginError = "";
      },
      err => {
        console.log(err);
        this.loginError = "Sorry, the username or password you entered is incorrect. Please try again.";
      }

    )
  }

}