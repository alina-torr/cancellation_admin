import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { number } from 'echarts';
import { Observable, catchError, of } from 'rxjs';
import { ApiUrl } from 'src/config';


@Injectable({
  providedIn: 'root'
})
export class BookingService {

  constructor(
    private http: HttpClient,
    // private errorService: ErrorService,
  ) { }

  getBookings(): Observable<any[] | null> {
    return this.http.get<any[]>(`${ApiUrl}/api/bookings`).pipe(
      catchError((error: HttpErrorResponse) => {
        return of(null);
      }),
    )
  }

  getStatistics(): Observable<Statistics | null> {
    return this.http.get<Statistics>(`${ApiUrl}/api/statistics`).pipe(
      catchError((error: HttpErrorResponse) => {
        return of(null);
      }),
    )
  }
}

export interface Statistics {
  DistributionChannel: ValueCount[];
  ProfitStat: {
    Future: {
      CanceledProfit: any[];
      NotCanceledProfit: any[];
    };
    Prev: {
      Profit: DateCount[];
    };
  }

};

interface ValueCount {
  Count: number;
  Value: string;
}

interface DateCount {
  Date: string;
  Value: number;
}

export interface TableBooking {
  CancellationPredict: number
  BookingId: number
  ArrivalDateYear: number
  ArrivalDateMonth: string
  ArrivalDateDayOfMonth: number
  StaysInWeekendNights: number
  StaysInWeekNights: number
  Adults: number
  Children: number
  Babies: number
  Meal: string
  RequiredCarParkingSpaces: number
}
