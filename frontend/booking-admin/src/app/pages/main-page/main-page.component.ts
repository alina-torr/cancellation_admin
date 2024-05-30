import { HttpClient } from '@angular/common/http';
import { OnInit } from '@angular/core';
import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';
import { EChartsOption } from 'echarts/types/dist/echarts';
import { AuthService } from 'src/app/services/auth/auth.service';
import { BookingService, Statistics, TableBooking } from 'src/app/services/bookings/booking.service';
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
      console.log(this.onlyLoadForModel)
    })
  }

  isLoading = false;
  hint = false;

  bookings!: any[];
  statistics!: Statistics;
  DCmergeOption: any;
  ProfitMergeOption: any;

  DCoptions: EChartsOption = {
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b} : {c} ({d}%)',
    },
    legend: {
      orient: 'vertical',
      left: 'left',
      data: ['USA', 'Germany', 'France', 'Canada', 'Russia'],

    },
    series: [
      {
        name: 'Distribution Channel',
        type: 'pie',
        radius: '80%',
        center: ['50%', '50%'],
      },
    ],
  };


  ProfitOptions: EChartsOption = {
    xAxis: {
      data: ['A', 'B', 'C', 'D', 'E']
    },
    yAxis: {},
    series: [
      {
        data: [10, 22, 28, 23, 19],
        type: 'line',
        smooth: true
      }
    ]
  };
  yearControl!: FormControl;
  monthControl!: FormControl;
  dayControl!: FormControl;

  tableData!: MatTableDataSource<TableBooking>;
  displayedColumns = [
    "CancellationPredict",
    "BookingId",
    "ArrivalDateYear",
    "ArrivalDateMonth",
    "ArrivalDateDayOfMonth",
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

  getMonth(i: number) {
    return this.months[i as keyof typeof this.months];
  }
  ngOnInit(): void {

    this.bookingService.getBookings().subscribe((res) => {
      if (res) {
        this.tableData = new MatTableDataSource(res);
        const yearSet = new Set();

        let minDate = new Date(res.reduce((curMin: number, d) => {
          if (Date.parse(d.Date) < curMin) {
            return Date.parse(`${d.ArrivalDateDayOfMonth} ${d.ArrivalDateMonth} ${d.ArrivalDateYear}`)
          }
          return curMin;
        }, Date.now()))

        this.yearControl = new FormControl(minDate.getFullYear());
        this.monthControl = new FormControl(minDate.getMonth());
        this.dayControl = new FormControl(minDate.getDate());
        // let maxDate = res.ProfitStat.Future.CanceledProfit.concat(res.ProfitStat.Future.NotCanceledProfit).reduce((curMax: number, d) => {
        //   if (Date.parse(d.Date) > curMax) {
        //     return Date.parse(d.Date)
        //   }
        //   return curMax;
        // }, Date.now())
      }
    });
    this.bookingService.getStatistics().subscribe((res) => {
      if (res) {
        this.statistics = res;
        this.DCmergeOption = {
          series: [
            { data: res.DistributionChannel.map((v) => ({
                value: v.Count,
                name: v.Value,
              }))
            },
          ],
          legend: {
            orient: 'vertical',
            left: 'left',
            data: res.DistributionChannel.map((v) => v.Value),
          },
        };
        console.log(res.ProfitStat.Prev)

        let minDate = res.ProfitStat.Prev.Profit.reduce((curMin: number, d) => {
          if (Date.parse(d.Date) < curMin) {
            return Date.parse(d.Date)
          }
          return curMin;
        }, Date.now())
        let maxDate = res.ProfitStat.Future.CanceledProfit.concat(res.ProfitStat.Future.NotCanceledProfit).reduce((curMax: number, d) => {
          if (Date.parse(d.Date) > curMax) {
            return Date.parse(d.Date)
          }
          return curMax;
        }, Date.now())

        let daysOfYear = [];
        let prev = [];
        let futureCanceled = [];
        let futureNotCanceled = [];
        for (var d = new Date(minDate); d <= new Date(maxDate); d.setDate(d.getDate() + 1)) {
            let k = new Date(d);
            daysOfYear.push(`${k.getDate()}.${k.getMonth()+1}.${k.getFullYear()}`);
            const month = {
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
            prev.push(res.ProfitStat.Prev.Profit.find((el) => {
              let mn = k.getMonth() as keyof typeof month
              return el.Date == `${k.getDate()} ${month[mn]} ${k.getFullYear()}`
            })?.Value)
            futureCanceled.push(res.ProfitStat.Future.CanceledProfit.find((el) => {
              let mn = k.getMonth() as keyof typeof month
              return el.Date == `${k.getDate()} ${month[mn]} ${k.getFullYear()}`
            })?.Value)
            futureNotCanceled.push(res.ProfitStat.Future.NotCanceledProfit.find((el) => {
              let mn = k.getMonth() as keyof typeof month
              return el.Date == `${k.getDate()} ${month[mn]} ${k.getFullYear()}`
            })?.Value)
        }
        console.log(daysOfYear)
        console.log(futureCanceled)
        console.log(futureNotCanceled)
        this.ProfitMergeOption = {
          xAxis: {
            data: daysOfYear
          },
          // series: [
          //   {
          //     data: futureCanceled
          //   },
          //   {
          //     data: prev
          //   },
          //   {
          //     data: futureNotCanceled
          //   },
          // ],
          series: [
            {
              name: 'Email',
              type: 'line',
              stack: 'Total',
              data: prev
            },
            {
              name: 'Union Ads',
              type: 'line',
              stack: 'Total',
              data: futureCanceled
            },
            {
              name: 'Video Ads',
              type: 'line',
              stack: 'Total',
              data: futureNotCanceled
            },
          ]
        };
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
          this.isLoading = false;
        });
    }
  }

  openHint() {
    this.hint = true;
  }

  onFileSelected(event: any) {
    const file: File = event.target.files[0];
    if (file) {
        const formData = new FormData();
        formData.append("file", file);
        const upload$ = this.http.post(`${ApiUrl}/api/upload_data`, formData);
        upload$.subscribe();
    }
  }

}

