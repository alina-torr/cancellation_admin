import { Component } from '@angular/core';
import { NonNullableFormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth/auth.service';
import { LoginResponse } from 'src/app/services/auth/models/auth';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.sass']
})
export class LoginPageComponent {
  errorMessage = '';

  form = this.formBuilder.group({
    login: ['', [Validators.required]],
    password: ['', [Validators.required]],
  });

  isLoading = false;

  constructor(
    private formBuilder: NonNullableFormBuilder,
    private authService: AuthService,
    private router: Router,
  ) {
    if (this.authService.getJwtToken()) {
      this.toMain();
    }
  }

  login() {
    if (this.form.valid && this.form.value.login && this.form.value.password) {
      this.isLoading = true;
      this.authService.login({login: this.form.value.login, password: this.form.value.password}).subscribe((res: LoginResponse | null) => {
        this.isLoading = false;
        if (res !== null) {
          this.authService.setJwtToken(res.access_token);
          this.toMain();
        }
      },
      (err) => {
        if (err.error.message === "2") {
          this.errorMessage = 'Неверный логин или пароль'
        }
        this.isLoading = false;
      })
    }
  }

  toMain() {
    this.router.navigate(['main']);
  }
}
