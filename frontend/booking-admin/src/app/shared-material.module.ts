import { NgModule } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import {MatSnackBarModule} from '@angular/material/snack-bar';

const modules: any[] = [
  MatIconModule,
  MatButtonModule,
  MatInputModule,
  MatProgressSpinnerModule,
  MatSnackBarModule,
];

@NgModule({
  declarations: [],
  imports: [modules],
  exports: [...modules],
})
export class SharedMaterialModule { }
