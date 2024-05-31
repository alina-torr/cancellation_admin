package repositories

import (
	ent "booking/entities"
	mlgrpc "booking/ml/grpc"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thoas/go-funk"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BookingRepository struct {
	dbpool *pgxpool.Pool
}

func NewBookingRepository(dbpool *pgxpool.Pool) *BookingRepository {
	return &BookingRepository{
		dbpool: dbpool,
	}
}

func (br BookingRepository) SaveBookingPredictions(bookings []ent.BookingFields, predictions []float32, hotelId int64) error {
	tx, err := br.dbpool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	for i, b := range bookings {
		// var bookingId int64
		_, err := tx.Exec(context.Background(),
			`INSERT INTO bookingscheme.booking (`+
				`hotel_id, booking_id, arrival_date_year, arrival_date_month, arrival_date_day_of_month, cancel_prediction, add_info`+
				`) VALUES (`+
				`$1, $2, $3, $4, $5, $6, $7`+
				`) ON CONFLICT (hotel_id, booking_id)
					DO UPDATE SET
					arrival_date_year = EXCLUDED.arrival_date_year,
					arrival_date_month = EXCLUDED.arrival_date_month,
					arrival_date_day_of_month = EXCLUDED.arrival_date_day_of_month,
					cancel_prediction = EXCLUDED.cancel_prediction,
					add_info = EXCLUDED.add_info;
					`,
			hotelId,
			b.BookingId,
			b.ArrivalDateYear,
			b.ArrivalDateMonth,
			b.ArrivalDateDayOfMonth,
			predictions[i],
			b.AddInfo,
		)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (br BookingRepository) GetAllForManager(managerId int64) ([]ent.BookingTable, error) {
	res, err := getArrayQuery[ent.BookingTable](br.dbpool,
		`select b.cancel_prediction, b.booking_id, b.arrival_date_year, b.arrival_date_month, b.arrival_date_day_of_month, b.add_info from userscheme.manager m 
			join hotelscheme.hotel h on m.hotel_id = h.id 
			join bookingscheme.booking b on h.id = b.hotel_id 
		where m.id = $1;`, managerId)
	return res, err
}

func (br BookingRepository) TrainModel(bookings []ent.BookingFields, cancellations []int64, hotelId int64) error {
	conn, err := grpc.Dial("localhost:50051",
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*1024),
			grpc.MaxCallSendMsgSize(1024*1024*1024),
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
			grpc.MaxCallRecvMsgSize(1024*1024*1024),
			grpc.MaxCallSendMsgSize(1024*1024*1024),
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

func String(v string) *string { return &v }
