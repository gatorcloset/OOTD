import { Component } from '@angular/core';
import {FormControl, Validators} from '@angular/forms';
import { LoginRequest } from '../mock-data/user';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent {
  formValid = true;
  username = new FormControl('', [Validators.required, Validators.pattern('\\S+')]);
  password = new FormControl('', [Validators.required]);
  loginError: string = "";

  constructor(private userService: UserService) {
    // subscribe to valueChanges of form controls to update the form validity
    this.username.valueChanges.subscribe(() => this.updateFormValidity());
    this.password.valueChanges.subscribe(() => this.updateFormValidity());
  }

  updateFormValidity() {
    this.formValid = this.username.invalid || this.password.invalid;
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
    this.updateFormValidity();
    this.userService.loginUser({ username, password } as LoginRequest).subscribe(
      res => {
        // Successful login => save authenticated user
        console.log(sessionStorage.getItem('userID'));

        // Test
        this.userService.setAuthUser(res);
        console.log(this.userService.getAuthUser().ID);

        // Reset error message
        this.loginError = "";
      },
      err => {
        console.log(err);
        this.loginError = "Sorry, the username or password you entered is incorrect. Please try again.";
      }

    )
  }

}