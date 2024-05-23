import { HttpClient } from '@angular/common/http';
import { OnInit } from '@angular/core';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import * as echarts from 'echarts/types/dist/echarts';
import { EChartsOption } from 'echarts/types/dist/echarts';
import { AuthService } from 'src/app/services/auth/auth.service';
import { BookingService, Statistics, TableBooking } from 'src/app/services/bookings/booking.service';
import { ApiUrl } from 'src/config';

@Component({
  selector: 'app-main-page',
  // standalone: true,
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.sass'],
  // providers: [
  //   provideEcharts(),
  // ]
})
export class MainPageComponent implements OnInit {

  constructor(
    private authService: AuthService,
    private bookingService: BookingService,
    private router: Router,
    private http: HttpClient,
  ) {
    if (!this.authService.getJwtToken()) {
      this.toLogin();
    }
  }

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

  tableData: TableBooking[] = [];
  displayedColumns = [
    "CancellationPredict",
    "BookingId",
    "ArrivalDateYear",
    "ArrivalDateMonth",
    "ArrivalDateDayOfMonth",
    "StaysInWeekendNights",
    "StaysInWeekNights",
    "Adults",
    "Children",
    "Babies",
    "Meal",
    "RequiredCarParkingSpaces",
  ]
  ngOnInit(): void {
    this.bookingService.getBookings().subscribe((res) => {
      if (res) {
        this.tableData = res;
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
        const upload$ = this.http.post(`${ApiUrl}/api/upload_data_predictions`, formData);
        upload$.subscribe();
    }
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

