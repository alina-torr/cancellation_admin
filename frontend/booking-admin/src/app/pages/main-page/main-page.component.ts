import { HttpClient } from '@angular/common/http';
import { OnInit } from '@angular/core';
import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';
import { EChartsOption } from 'echarts/types/dist/echarts';
import { AuthService } from 'src/app/services/auth/auth.service';
import { BookingService, TableBooking } from 'src/app/services/bookings/booking.service';
import { ApiUrl } from 'src/config';

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.sass'],
})
export class MainPageComponent implements OnInit {
  onlyLoadForModel: boolean | undefined;
  constructor(
    private authService: AuthService,
    private bookingService: BookingService,
    private router: Router,
    private http: HttpClient,
    public dialog: MatDialog,
  ) {
    if (!this.authService.getJwtToken()) {
      this.toLogin();
    }
    this.bookingService.getIsThereModel().subscribe((res: boolean) => {
      this.onlyLoadForModel = !res;
      if (res) this.getBookings();
    })
  }

  isLoading = false;
  isError = false;
  hint = false;

  bookings!: any[];

  yearControl!: FormControl;
  monthControl!: FormControl;
  dayControl!: FormControl;

  tableData!: MatTableDataSource<TableBooking>;
  displayedColumns = [
    "BookingId",
    "ArrivalDateYear",
    "ArrivalDateMonth",
    "ArrivalDateDayOfMonth",
    "CancellationPredict",
    "AddInfo"
  ]

  months = {
    0: 'January',
    1: 'February',
    2: 'March',
    3: 'April',
    4: 'May',
    5: 'June',
    6: 'July',
    7: 'August',
    8: 'September',
    9: 'October',
    10: 'November',
    11: 'December',
  }

  years: any[] = [];

  getMonth(i: number) {
    return this.months[i as keyof typeof this.months];
  }

  nestedRes: any = {};

  ngOnInit(): void {
    this.getBookings();
  }
  isDataReady = false;
  getBookings() {
    this.isLoading = true;
    this.bookingService.getBookings().subscribe((res) => {
      if (res) {
        this.tableData = new MatTableDataSource(res);
        this.isDataReady = true;
        const yearSet = new Set();
        let date = new Date();
        date.setDate(date.getDate() + 365);
        let minDate = new Date(res.reduce((curMin: number, d) => {
          yearSet.add(d.ArrivalDateYear)
          if (!this.nestedRes[d.ArrivalDateYear]) {
            this.nestedRes[d.ArrivalDateYear] = {};
          }
          if (!this.nestedRes[d.ArrivalDateYear][d.ArrivalDateMonth]) {
            this.nestedRes[d.ArrivalDateYear][d.ArrivalDateMonth] = {};
          }
          if (!this.nestedRes[d.ArrivalDateYear][d.ArrivalDateMonth][d.ArrivalDateDayOfMonth]) {
            this.nestedRes[d.ArrivalDateYear][d.ArrivalDateMonth][d.ArrivalDateDayOfMonth] = [];
          }
          this.nestedRes[d.ArrivalDateYear][d.ArrivalDateMonth][d.ArrivalDateDayOfMonth].push(d);
          if (Date.parse(`${d.ArrivalDateDayOfMonth} ${d.ArrivalDateMonth} ${d.ArrivalDateYear}`) < curMin) {
            return Date.parse(`${d.ArrivalDateDayOfMonth} ${d.ArrivalDateMonth} ${d.ArrivalDateYear}`)
          }
          return curMin;
        }, date.getTime()))
        this.years = Array.from(yearSet)
        this.yearControl = new FormControl(minDate.getFullYear());
        this.monthControl = new FormControl(minDate.getMonth());
        this.dayControl = new FormControl(minDate.getDate() - 1);


        const getFiltered = () => {
          let d = {
            ArrivalDateYear: this.yearControl.value,
            ArrivalDateMonth: this.months[this.monthControl.value as keyof typeof this.months],
            ArrivalDateDayOfMonth: this.dayControl.value + 1,
          }
          if (this.nestedRes[d.ArrivalDateYear] && this.nestedRes[d.ArrivalDateYear][d.ArrivalDateMonth] && this.nestedRes[d.ArrivalDateYear][d.ArrivalDateMonth][d.ArrivalDateDayOfMonth]) {
            return (this.nestedRes[d.ArrivalDateYear][d.ArrivalDateMonth][d.ArrivalDateDayOfMonth])
          }
          return [];
        }

        this.tableData = new MatTableDataSource(getFiltered());

        this.yearControl.valueChanges.subscribe(() => {
          this.tableData.data = getFiltered();
        })
        this.monthControl.valueChanges.subscribe(() => {
          this.tableData.data = getFiltered();
        })
        this.dayControl.valueChanges.subscribe(() => {
          this.tableData.data = getFiltered();
        })
        this.isLoading = false;
      }
    });
  }

  toLogin() {
    this.router.navigate(['login']);
  }

  onGetPredictions(event: any) {
    const file: File = event.target.files[0];
    if (file) {
        const formData = new FormData();
        formData.append("file", file);
        this.isLoading = true;
        const upload$ = this.http.post(`${ApiUrl}/api/upload_data_predictions`, formData);
        upload$.subscribe(() => {
          this.getBookings();
        },
        (error) => {
          this.isLoading = false;
          this.isError = true;
        });
    }
  }

  openHint() {
    this.hint = true;
  }

  onFileSelected(event: any, isStart = false) {
    const file: File = event.target.files[0];
    if (file) {
        const formData = new FormData();
        formData.append("file", file);
        const upload$ = this.http.post(`${ApiUrl}/api/upload_data`, formData);
        this.isLoading = true;
        upload$.subscribe(() => {
          this.isLoading = false;
          if (isStart) {
            this.router.navigateByUrl('/', {skipLocationChange: true}).then(() => {
              this.router.navigate(['/main']);
          });
          }
        }, (error) => {
          this.isLoading = false;
          this.isError = true;
        });
    }
  }

}

