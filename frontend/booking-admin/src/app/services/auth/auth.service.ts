import { Injectable } from '@angular/core';
import { LoginData, LoginResponse, RegisterData } from './models/auth';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, catchError, of } from 'rxjs';
import { ApiUrl } from 'src/config';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private jwtToken?: string;

  constructor(
    private http: HttpClient,
    private router: Router,
    // private errorService: ErrorService,
  ) {
    let jwt = localStorage.getItem('jwt');
    if (jwt) {
      this.jwtToken = jwt;
    }
  }

  setJwtToken(token: string) {
    localStorage.setItem('jwt', token);
    this.jwtToken = token;
  }

  getJwtToken(): string | undefined {
    return this.jwtToken;
  }

  login(loginData: LoginData): Observable<LoginResponse | null> {
    return this.http.post<LoginResponse>(`${ApiUrl}/login`, loginData);
  }

  register(loginData: RegisterData): Observable<LoginResponse | null> {
    return this.http.post<LoginResponse>(`${ApiUrl}/register`, loginData);
  }

  logout() {
    this.jwtToken = undefined;
    localStorage.removeItem('jwt')
    this.router.navigate(['login']);
  }

}
