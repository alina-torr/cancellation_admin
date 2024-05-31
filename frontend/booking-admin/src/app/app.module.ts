import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginPageComponent } from './pages/login-page/login-page.component';
import { SharedMaterialModule } from './shared-material.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { ReactiveFormsModule } from '@angular/forms';
import { JwtAuthInterceptor } from './services/auth/middleware/jwt-middleware';
import { NgxEchartsModule } from 'ngx-echarts';
import { MainPageComponent } from './pages/main-page/main-page.component';
import {MatTableModule} from '@angular/material/table';
import { RegisterPageComponent } from './pages/register-page/register-page.component';
import { TopBarComponent } from './components/top-bar/top-bar.component';
import {MatDialogModule} from '@angular/material/dialog';
import { LoadingDialogComponent } from './components/loading-dialog/loading-dialog.component';
import { FormatHintComponent } from './components/format-hint/format-hint.component';
import {MatSelectModule} from '@angular/material/select';
import {MatFormFieldModule} from '@angular/material/form-field';
import { ErrorComponent } from './components/error/error.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginPageComponent,
    RegisterPageComponent,
    MainPageComponent,
    TopBarComponent,
    LoadingDialogComponent,
    FormatHintComponent,
    ErrorComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    SharedMaterialModule,
    ReactiveFormsModule,
    HttpClientModule,
    MatDialogModule,
    NgxEchartsModule.forRoot({
      echarts: () => import('echarts'),
    }),
    MatTableModule,
    MatFormFieldModule,
    MatSelectModule,
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: JwtAuthInterceptor,
      multi: true,
    },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
