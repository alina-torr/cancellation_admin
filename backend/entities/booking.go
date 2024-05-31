package entities

type Booking struct {
	BookingFields
}

type TrainBooking struct {
	IsCanceled int64 `csv:"is_canceled"`
	BookingFields
}

type BookingFields struct {
	BookingId                   int64   `csv:"booking_id"`
	Leadtime                    int64   `csv:"lead_time"`
	ArrivalDateYear             int64   `csv:"arrival_date_year"`
	ArrivalDateMonth            string  `csv:"arrival_date_month"`
	ArrivalDateWeekNumber       int64   `csv:"arrival_date_week_number"`
	ArrivalDateDayOfMonth       int64   `csv:"arrival_date_day_of_month"`
	StaysInWeekendNights        int64   `csv:"stays_in_weekend_nights"`
	StaysInWeekNights           int64   `csv:"stays_in_week_nights"`
	Adults                      int64   `csv:"adults"`
	Children                    int64   `csv:"children"`
	Babies                      int64   `csv:"babies"`
	Meal                        string  `csv:"meal"`
	Country                     string  `csv:"country"`
	MarketSegment               string  `csv:"market_segment"`
	DistributionChannel         string  `csv:"distribution_channel"`
	PreviousCancellations       int64   `csv:"previous_cancellations"`
	PreviousBookingsNotCanceled int64   `csv:"previous_bookings_not_canceled"`
	ReservedRoomType            string  `csv:"reserved_room_type"`
	AssignedRoomType            string  `csv:"assigned_room_type"`
	BookingChanges              int64   `csv:"booking_changes"`
	Agent                       string  `csv:"agent"`
	Company                     string  `csv:"company"`
	DaysInWaitingList           int64   `csv:"days_in_waiting_list"`
	CustomerType                string  `csv:"customer_type"`
	Adr                         float64 `csv:"adr"`
	RequiredCarParkingSpaces    int64   `csv:"required_car_parking_spaces"`
	TotalOfSpecialRequests      int64   `csv:"total_of_special_requests"`
	AddInfo                     string  `csv:"add_info"`
}

type BookingTable struct {
	CancellationPredict   float32 `db:"cancel_prediction"`
	BookingId             int64   `db:"booking_id"`
	ArrivalDateYear       int64   `db:"arrival_date_year"`
	ArrivalDateMonth      string  `db:"arrival_date_month"`
	ArrivalDateDayOfMonth int64   `db:"arrival_date_day_of_month"`
	AddInfo               string  `db:"add_info"`
}

func CastBookingFieldsToTable(b BookingFields, predict float32) BookingTable {
	return BookingTable{
		CancellationPredict:   predict,
		BookingId:             b.BookingId,
		ArrivalDateYear:       b.ArrivalDateYear,
		ArrivalDateMonth:      b.ArrivalDateMonth,
		ArrivalDateDayOfMonth: b.ArrivalDateDayOfMonth,
		AddInfo:               b.AddInfo,
	}
}
