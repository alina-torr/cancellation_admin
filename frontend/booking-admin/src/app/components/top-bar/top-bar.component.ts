import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth/auth.service';

@Component({
  selector: 'app-top-bar',
  templateUrl: './top-bar.component.html',
  styleUrls: ['./top-bar.component.sass']
})
export class TopBarComponent {

  constructor(private authService: AuthService, private router: Router) {}

  logout() {
    this.authService.logout();
  }

  navigateLogin() {
    this.router.navigate(['login']);
  }

  navigateRegister() {
    this.router.navigate(['register']);
  }

  isLogged() {
    return this.authService.getJwtToken()
  }
}
