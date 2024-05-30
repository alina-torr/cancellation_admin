package entities

type Booking struct {
	Id int64 `db:"id"`
	BookingFields
}

// type BookingFields struct {
// 	BookingId                   int    `db:"booking_id"`
// 	IsCanceled                  int64  `db:"is_canceled"`
// 	Leadtime                    int    `db:"lead_time"`
// 	ArrivalDateYear             int    `db:"arrival_date_year"`
// 	ArrivalDateMonth            string `db:"arrival_date_month"`
// 	ArrivalDateWeekNumber       int    `db:"arrival_date_week_number"`
// 	ArrivalDateDayOfMonth       int    `db:"arrival_date_day_of_month"`
// 	StaysInWeekendNights        int    `db:"stays_in_weekend_nights"`
// 	StaysInWeekNights           int    `db:"stays_in_week_nights"`
// 	AdultNumber                 int    `db:"adult_number"`
// 	ChildNumber                 int    `db:"child_number"`
// 	BabyNumber                  int    `db:"baby_number"`
// 	Meal                        string `db:"meal"`
// 	Country                     string `db:"country"`
// 	MarketSegment               string `db:"market_segment"`
// 	DistributionChannel         string `db:"distribution_channel"`
// 	IsRepeatedGuest             bool   `db:"is_repeated_guest"`
// 	PreviousCancellations       int    `db:"previous_cancellations"`
// 	PreviousBookingsNotCanceled int    `db:"previous_bookings_not_canceled"`
// 	ReservedRoomType            string `db:"reserved_room_type"`
// 	AssignedRoomType            string `db:"assigned_room_type"`
// 	BookingChanges              int    `db:"booking_changes"`
// 	DepositType                 string `db:"deposit_type"`
// 	Agent                       string `db:"agent"`
// 	Company                     string `db:"company"`
// 	DaysInWaitingList           int    `db:"days_in_waiting_list"`
// 	CustomerType                string `db:"customer_type"`
// 	Adr                         int    `db:"adr"`
// 	RequiredCarParkingApaces    int    `db:"required_car_parking_spaces"`
// 	TotalOfSpecialRequests      int    `db:"total_of_special_requests"`
// }

type TrainBooking struct {
	IsCanceled int64 `csv:"is_canceled"`
	BookingFields
}

type BookingFields struct {
	BookingId             int64  `csv:"booking_id"`
	Leadtime              int64  `csv:"lead_time"`
	ArrivalDateYear       int64  `csv:"arrival_date_year"`
	ArrivalDateMonth      string `csv:"arrival_date_month"`
	ArrivalDateWeekNumber int64  `csv:"arrival_date_week_number"`
	ArrivalDateDayOfMonth int64  `csv:"arrival_date_day_of_month"`
	StaysInWeekendNights  int64  `csv:"stays_in_weekend_nights"`
	StaysInWeekNights     int64  `csv:"stays_in_week_nights"`
	Adults                int64  `csv:"adults"`
	Children              int64  `csv:"children"`
	Babies                int64  `csv:"babies"`
	Meal                  string `csv:"meal"`
	Country               string `csv:"country"`
	MarketSegment         string `csv:"market_segment"`
	DistributionChannel   string `csv:"distribution_channel"`
	// IsRepeatedGuest             bool    `csv:"is_repeated_guest"`
	PreviousCancellations       int64  `csv:"previous_cancellations"`
	PreviousBookingsNotCanceled int64  `csv:"previous_bookings_not_canceled"`
	ReservedRoomType            string `csv:"reserved_room_type"`
	AssignedRoomType            string `csv:"assigned_room_type"`
	BookingChanges              int64  `csv:"booking_changes"`
	// DepositType                 string  `csv:"deposit_type"`
	Agent                    string  `csv:"agent"`
	Company                  string  `csv:"company"`
	DaysInWaitingList        int64   `csv:"days_in_waiting_list"`
	CustomerType             string  `csv:"customer_type"`
	Adr                      float64 `csv:"adr"`
	RequiredCarParkingSpaces int64   `csv:"required_car_parking_spaces"`
	TotalOfSpecialRequests   int64   `csv:"total_of_special_requests"`
}

type BookingTable struct {
	CancellationPredict   float32 `db:"cancel_prediction"`
	BookingId             int64   `db:"booking_id"`
	ArrivalDateYear       int64   `db:"arrival_date_year"`
	ArrivalDateMonth      string  `db:"arrival_date_month"`
	ArrivalDateDayOfMonth int64   `db:"arrival_date_day_of_month"`
	// StaysInWeekendNights     int64
	// StaysInWeekNights        int64
	// Adults                   int64
	// Children                 int64
	// Babies                   int64
	// Meal                     string
	// RequiredCarParkingSpaces int64
}

// func CastCSVtoDB(b BookingCSV, hotelId int) BookingFields {
// 	return BookingFields{
// 		BookingId:                   b.BookingId,
// 		IsCanceled:                  b.IsCanceled,
// 		Leadtime:                    b.Leadtime,
// 		ArrivalDateYear:             b.ArrivalDateYear,
// 		ArrivalDateMonth:            b.ArrivalDateMonth,
// 		ArrivalDateWeekNumber:       b.ArrivalDateWeekNumber,
// 		ArrivalDateDayOfMonth:       b.ArrivalDateDayOfMonth,
// 		StaysInWeekendNights:        b.StaysInWeekendNights,
// 		StaysInWeekNights:           b.StaysInWeekNights,
// 		AdultNumber:                 b.Adults,
// 		ChildNumber:                 b.Children,
// 		BabyNumber:                  b.Babies,
// 		Meal:                        b.Meal,
// 		Country:                     b.Country,
// 		MarketSegment:               b.MarketSegment,
// 		DistributionChannel:         b.DistributionChannel,
// 		IsRepeatedGuest:             b.IsRepeatedGuest,
// 		PreviousCancellations:       b.PreviousCancellations,
// 		PreviousBookingsNotCanceled: b.PreviousBookingsNotCanceled,
// 		ReservedRoomType:            b.ReservedRoomType,
// 		AssignedRoomType:            b.AssignedRoomType,
// 		BookingChanges:              b.BookingChanges,
// 		DepositType:                 b.DepositType,
// 		Agent:                       b.Agent,
// 		Company:                     b.Company,
// 		DaysInWaitingList:           b.DaysInWaitingList,
// 		CustomerType:                b.CustomerType,
// 		Adr:                         int(b.Adr),
// 		RequiredCarParkingApaces:    b.RequiredCarParkingSpaces,
// 		TotalOfSpecialRequests:      b.TotalOfSpecialRequests,
// 	}
// }

func CastBookingFieldsToTable(b BookingFields, predict float32) BookingTable {
	return BookingTable{
		CancellationPredict:   predict,
		BookingId:             b.BookingId,
		ArrivalDateYear:       b.ArrivalDateYear,
		ArrivalDateMonth:      b.ArrivalDateMonth,
		ArrivalDateDayOfMonth: b.ArrivalDateDayOfMonth,
		// StaysInWeekendNights:     b.StaysInWeekendNights,
		// StaysInWeekNights:        b.StaysInWeekNights,
		// Adults:                   b.Adults,
		// Children:                 b.Children,
		// Babies:                   b.Babies,
		// Meal:                     b.Meal,
		// RequiredCarParkingSpaces: b.RequiredCarParkingSpaces,
	}
}
