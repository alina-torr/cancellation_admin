import { Component, EventEmitter, Inject, Output } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-format-hint',
  templateUrl: './format-hint.component.html',
  styleUrls: ['./format-hint.component.sass']
})
export class FormatHintComponent {
  @Output() close = new EventEmitter();
}
