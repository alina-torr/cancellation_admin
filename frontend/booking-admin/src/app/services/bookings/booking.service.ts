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
  ) { }

  getBookings(): Observable<any[] | null> {
    return this.http.get<any[]>(`${ApiUrl}/api/get_predictions`).pipe(
      catchError((error: HttpErrorResponse) => {
        return of(null);
      }),
    )
  }

  getIsThereModel(): Observable<boolean> {
    return this.http.get<boolean>(`${ApiUrl}/api/is_there_model`);
  }

}

export interface TableBooking {
  CancellationPredict: number
  BookingId: number
  ArrivalDateYear: number
  ArrivalDateMonth: string
  ArrivalDateDayOfMonth: number
}
