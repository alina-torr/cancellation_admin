package repositories

import (
	ent "booking/entities"
	"context"
	"fmt"
	"math"

	// "fmt"

	// "fmt"
	mlgrpc "booking/ml/grpc"
	// "github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx"
	// "github.com/jackc/pgx/v5"
	"github.com/thoas/go-funk"

	// "github.com/thoas/go-funk"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// "github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingRepository struct {
	dbpool *pgxpool.Pool
}

func NewBookingRepository(dbpool *pgxpool.Pool) *BookingRepository {
	return &BookingRepository{
		dbpool: dbpool,
	}
}

// func (hr HotelRepository) GetById(id int, managerId int) (ent.Hotel, error) {
// 	return ent.Hotel{}, nil
// }

// func (br BookingRepository) CreateInTx(b ent.BookingFields, tx pgx.Tx) (int64, error) {
// 	var bookingId int64
// 	// err = tx.QueryRow(context.Background(),
// 	// 	`INSERT INTO bookingscheme.booking (`+
// 	// 		`client_id, hotel_id, reserved_room_id, assigned_room_id, night_stays, adult_number, child_number, baby_number, special_requests, agent, company, required_car_parking_spaces, arrival_date, booking_date, confirmed_date, is_canceled, meal, market_segment, distribution_channel, customer_type, deposit, booking_changes, reservation_status`+
// 	// 		`) VALUES (`+
// 	// 		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23`+
// 	// 		`) RETURNING id`,
// 	// 	b.ClientId,
// 	// 	b.HotelId,
// 	// 	b.ReservedRoomId,
// 	// 	b.AssignedRoomId,
// 	// 	b.NightStays,
// 	// 	b.Adults,
// 	// 	b.Children,
// 	// 	b.Babies,
// 	// 	b.SpecialRequests,
// 	// 	b.Agent,
// 	// 	b.Company,
// 	// 	b.RequiredCarParkingSpaces,
// 	// 	b.ArrivalDate,
// 	// 	b.BookingDate,
// 	// 	b.ConfirmedDate,
// 	// 	b.IsCanceled,
// 	// 	b.Meal,
// 	// 	b.MarketSegment,
// 	// 	b.DistributionChannel,
// 	// 	b.CustomerType,
// 	// 	b.Deposit,
// 	// 	b.BookingChanges,
// 	// 	b.ReservationStatus).Scan(&bookingId)
// 	err := tx.QueryRow(context.Background(),
// 		`INSERT INTO bookingscheme.booking (`+
// 			`hotel_id, booking_id, is_canceled, lead_time, arrival_date_year, arrival_date_month, arrival_date_week_number, arrival_date_day_of_month, stays_in_weekend_nights, stays_in_week_nights, adult_number,child_number,baby_number, meal, country, market_segment,distribution_channel,is_repeated_guest, previous_cancellations, previous_bookings_not_canceled, reserved_room_type, assigned_room_type, booking_changes,deposit_type,agent,company,days_in_waiting_list,customer_type,adr,required_car_parking_spaces,total_of_special_requests,reservation_status`+
// 			`) VALUES (`+
// 			`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32`+
// 			`) RETURNING id`,
// 		b.HotelId,
// 		b.BookingId,
// 		b.IsCanceled,
// 		b.Leadtime,
// 		b.ArrivalDateYear,
// 		b.ArrivalDateMonth,
// 		b.ArrivalDateWeekNumber,
// 		b.ArrivalDateDayOfMonth,
// 		b.StaysInWeekendNights,
// 		b.StaysInWeekNights,
// 		b.AdultNumber,
// 		b.ChildNumber,
// 		b.BabyNumber,
// 		b.Meal,
// 		b.Country,
// 		b.MarketSegment,
// 		b.DistributionChannel,
// 		b.IsRepeatedGuest,
// 		b.PreviousCancellations,
// 		b.PreviousBookingsNotCanceled,
// 		b.ReservedRoomType,
// 		b.AssignedRoomType,
// 		b.BookingChanges,
// 		b.DepositType,
// 		b.Agent,
// 		b.Company,
// 		b.DaysInWaitingList,
// 		b.CustomerType,
// 		b.Adr,
// 		b.RequiredCarParkingApaces,
// 		b.TotalOfSpecialRequests,
// 		b.ReservationStatus).Scan(&bookingId)

// 	if err != nil {
// 		return -1, err
// 	}

// 	return bookingId, nil
// }

// func (br BookingRepository) Create(b ent.BookingFields) (int64, error) {
// 	tx, err := br.dbpool.Begin(context.Background())
// 	if err != nil {
// 		return -1, err
// 	}
// 	defer tx.Rollback(context.Background())

// 	bookingId, err := br.CreateInTx(b, tx)

// 	if err != nil {
// 		return -1, err
// 	}

// 	err = tx.Commit(context.Background())
// 	if err != nil {
// 		return -1, err
// 	}
// 	return bookingId, nil
// }

// func (br BookingRepository) CreateBookings(bs []ent.BookingFields) error {
// 	tx, err := br.dbpool.Begin(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback(context.Background())

// 	for _, b := range bs {
// 		_, err := br.CreateInTx(b, tx)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit(context.Background())
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (br BookingRepository) GetAllForManager(managerId int64) ([]ent.Booking, error) {
	res, err := getArrayQuery[ent.Booking](br.dbpool,
		`select b.* from userscheme.manager m 
			join hotelscheme.hotel h on m.hotel_id = h.id 
			join bookingscheme.booking b on h.id = b.hotel_id 
		where m.id = $1;`, managerId)
	return res, err
}

func (br BookingRepository) GetDistributionChannel(managerId int64) ([]CountStatictic, error) {
	res, err := getArrayQuery[CountStatictic](br.dbpool,
		`select b.distribution_channel as value, count(b.distribution_channel) as count
			from userscheme.manager m 
				join hotelscheme.hotel h on m.hotel_id = h.id 
				join bookingscheme.booking b on h.id = b.hotel_id 
		where m.id=$1
		group by b.distribution_channel;`, managerId)
	return res, err
}

func (br BookingRepository) TrainModel(bookings []ent.BookingFields, cancellations []int64, hotelId int64) error {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*32),
			grpc.MaxCallSendMsgSize(math.MaxInt64),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := mlgrpc.NewMlClient(conn)
	bs := funk.Map(bookings, func(b ent.BookingFields) *mlgrpc.Booking {
		return br.BookingFieldsToMlBooking(b)
	}).([]*mlgrpc.Booking)

	res, err := client.TrainModel(context.Background(), &mlgrpc.BookingTrainRequest{Bookings: bs, IsCanceled: cancellations, HotelId: hotelId})
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (br BookingRepository) GetPredictions(bs []ent.BookingFields, hotelId int64) ([]float32, error) {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*32),
			grpc.MaxCallSendMsgSize(math.MaxInt64),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return make([]float32, 0), err
	}
	defer conn.Close()

	client := mlgrpc.NewMlClient(conn)
	bookings := funk.Map(bs, func(b ent.BookingFields) *mlgrpc.Booking {
		return br.BookingFieldsToMlBooking(b)
	}).([]*mlgrpc.Booking)
	if len(bookings) == 0 {
		return make([]float32, 0), nil
	}
	message, err := client.GetPredictions(context.Background(), &mlgrpc.BookingPredictRequest{Bookings: bookings, HotelId: hotelId})
	if err != nil {
		return make([]float32, 0), err
	}
	return message.Predictions, nil
}

func (br BookingRepository) BookingFieldsToMlBooking(b ent.BookingFields) *mlgrpc.Booking {
	return &mlgrpc.Booking{
		Leadtime:                    b.Leadtime,
		ArrivalDateMonth:            b.ArrivalDateMonth,
		ArrivalDateWeekNumber:       b.ArrivalDateWeekNumber,
		ArrivalDayOfMonth:           b.ArrivalDateDayOfMonth,
		StaysInWeekendNights:        b.StaysInWeekendNights,
		StaysInWeekNights:           b.StaysInWeekNights,
		Adults:                      b.Adults,
		Children:                    b.Children,
		Babies:                      b.Babies,
		Meal:                        &b.Meal,
		Country:                     &b.Country,
		MarketSegment:               b.MarketSegment,
		DistributionChannel:         &b.DistributionChannel,
		PreviousCancellations:       b.PreviousCancellations,
		PreviousBookingsNotCanceled: b.PreviousBookingsNotCanceled,
		ReservedRoomType:            b.ReservedRoomType,
		AssignedRoomType:            b.AssignedRoomType,
		BookingChanges:              b.BookingChanges,
		Agent:                       &b.Agent,
		Company:                     &b.Company,
		DaysInWaitingList:           b.DaysInWaitingList,
		CustomerType:                b.CustomerType,
		Adr:                         b.Adr,
		RequiredCarParkingSpaces:    b.RequiredCarParkingSpaces,
		TotalOfSpecialRequests:      b.TotalOfSpecialRequests,
	}
}

type CountStatictic struct {
	Value string `db:"value"`
	Count int    `db:"count"`
}

func String(v string) *string { return &v }
