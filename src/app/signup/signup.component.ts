import { Component } from '@angular/core';
import { User } from '../mock-data/user';
import { UserService } from '../services/user.service';
import {FormControl, Validators} from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent {
  newUser: User = <User>{};

  formValid = true;
  first = new FormControl('', [Validators.required, Validators.pattern('[a-zA-Z]+')]);
  last = new FormControl('', [Validators.required, Validators.pattern('[a-zA-Z]+')]);
  usernameControl = new FormControl('', [Validators.required, Validators.pattern('\\S+')]);
  passwordControl = new FormControl('', [Validators.required]);
  // confirmPassword = new FormControl('', [Validators.required]);
  signUpError: string = "";

  constructor(private userService: UserService, private router: Router) { 
    this.usernameControl.valueChanges.subscribe(() => this.updateFormValidity());
    this.passwordControl.valueChanges.subscribe(() => this.updateFormValidity());
    this.first.valueChanges.subscribe(() => this.updateFormValidity());
    this.last.valueChanges.subscribe(() => this.updateFormValidity());
    // this.confirmPassword.valueChanges.subscribe(() => this.updateFormValidity());
  }

  updateFormValidity() {
    this.formValid = this.usernameControl.invalid || this.passwordControl.invalid || this.first.invalid || this.last.invalid;
  }

  firstErrorMessage() {
    if (this.first.hasError('required')) {
      return 'You must enter a first name';
    }
    if (this.first.hasError('pattern')) {
      return 'Please enter a valid first name'
    }

    return '';
  }

  lastErrorMessage() {
    if (this.last.hasError('required')) {
      return 'You must enter a last name';
    }
    if (this.last.hasError('pattern')) {
      return 'Please enter a valid last name'
    }

    return '';
  }

  userErrorMessage() {
    if (this.usernameControl.hasError('required')) {
      return 'You must enter a username';
    }

    if (this.usernameControl.hasError('pattern')) {
      return 'Not a valid username'
    }
    
    return '';
  }

  passErrorMessage() {
    if (this.passwordControl.hasError('required')) {
      return 'You must enter a password'
    }
    
    return '';
  }

  /*
  confirmErrorMessage() {
    if (this.confirmPassword.hasError('required')) {
      return 'You must confirm your password';
    }

    if (this.passwordControl.value !== this.confirmPassword.value) {
      this.confirmPassword.setErrors({ 'notSame': true });
      return 'Your password must match';
    }
    else {
      this.confirmPassword.setErrors(null); // clear passwordMismatch error
    }
    

    return '';
  }
  */

  createUser(firstname: string, lastname: string, username: string, password: string) {
    this.updateFormValidity();

    // Trim any added spaces after the name
    firstname = firstname.trim();
    lastname = lastname.trim();
    username = username.trim();

    this.userService.createUser({ firstname, lastname, username, password } as User).subscribe(
      newUser => {
        this.newUser = newUser
        this.userService.authenticated = true;
        this.userService.authUser = newUser;
        this.router.navigateByUrl('/closet');
      },
      err => {
        console.log(err);
        this.userService.authenticated = false;
        this.userService.authUser = undefined;
      }
    )

  }
}
