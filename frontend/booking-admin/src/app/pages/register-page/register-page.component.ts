import { Component } from '@angular/core';
import { Validators, NonNullableFormBuilder } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth/auth.service';
import { LoginResponse } from 'src/app/services/auth/models/auth';

@Component({
  selector: 'app-register-page',
  templateUrl: './register-page.component.html',
  styleUrls: ['./register-page.component.sass']
})
export class RegisterPageComponent {
  errorMessage = '';
  form = this.formBuilder.group({
    login: ['', [Validators.required]],
    password: ['', [Validators.required]],
    hotelName: ['', [Validators.required]],
    hotelCity: ['', [Validators.required]],
    hotelCountry: ['', [Validators.required]],
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

  register() {
    if (this.form.valid && this.form.value.login && this.form.value.password && this.form.value.hotelName && this.form.value.hotelCity && this.form.value.hotelCountry) {
      this.isLoading = true;
      this.errorMessage = '';
      this.authService.register({
        login: this.form.value.login,
        password: this.form.value.password,
        hotelName: this.form.value.hotelName,
        hotelCity: this.form.value.hotelCity,
        hotelCountry: this.form.value.hotelCountry,

      }).subscribe((res: LoginResponse | null) => {
        this.isLoading = false;
        if (res !== null) {
          this.authService.setJwtToken(res.access_token);
          this.toMain();
        }
      },
      (err) => {
        if (err.error.message === "1") {
          this.errorMessage = 'Данный логин уже используется'
        }
        this.isLoading = false;
      }
      )
    }
  }

  toMain() {
    this.router.navigate(['main']);
  }
}
