import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, Router, createUrlTreeFromSnapshot } from '@angular/router';
import { UserService } from './user.service';

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private userService: UserService, private router: Router) {}

  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot) {

    // Check if user is authenticated
    if (this.userService.authenticated === true) {
      return true; // Allow access to the route
    } 
    else {
      this.router.navigateByUrl('/login');
      return false; // Prevent access to the route
    }
  }
}
