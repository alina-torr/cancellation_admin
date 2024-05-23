package services

import (
	ent "booking/entities"
	rep "booking/repositories"
	"fmt"
	// "sort"
	"strconv"
	"strings"
	"time"

	// "fmt"
	// "os"
	// "os/exec"
	// "time"
	mlgrpc "booking/ml/grpc"
	// "github.com/mitchellh/mapstructure"
	// "github.com/thoas/go-funk"
)

type bookingRepository interface {
	// Create(booking ent.BookingFields) (int64, error)
	// CreateBookings(bs []ent.BookingFields) error
	GetAllForManager(managerId int64) ([]ent.Booking, error)
	GetDistributionChannel(managerId int64) ([]rep.CountStatictic, error)
	GetPredictions(bs []ent.BookingFields, hotelId int64) ([]float32, error)
	TrainModel(bookings []ent.BookingFields, cancellations []int64, hotelId int64) error
}

type BookingService struct {
	bookingRepository bookingRepository
	userRepository    userRepository
}

func NewBookingService(bRepository bookingRepository, uRepository userRepository) *BookingService {
	return &BookingService{
		bookingRepository: bRepository,
		userRepository:    uRepository,
	}
}

// func (bs BookingService) Create(w ent.BookingFields) (int64, error) {
// 	return bs.bookingRepository.Create(w)
// }

// func (bs BookingService) GetAllForManager(managerId int64) ([]ent.Booking, error) {
// 	return bs.bookingRepository.GetAllForManager(managerId)
// }

// func (bs BookingService) GetDistributionChannel(managerId int64) ([]rep.CountStatictic, error) {
// 	return bs.bookingRepository.GetDistributionChannel(managerId)
// }

func (bs BookingService) GetPredicts(bookings []*mlgrpc.Booking) {
	res := mlgrpc.BookingPredictRequest{
		Bookings: bookings,
	}
	fmt.Println(res)
}

// func (bs BookingService) CreateBookings(bss []ent.BookingFields) error {
// 	return bs.bookingRepository.CreateBookings(bss)
// }

func (bs BookingService) GetPredictions(bss []ent.BookingFields, managerId int64) ([]float32, error) {
	m, err := bs.userRepository.GetManagerById(managerId)
	if err != nil {
		return make([]float32, 0), err
	}
	return bs.bookingRepository.GetPredictions(bss, m.HotelId)
}

func (bs BookingService) GetPrediction(bss ent.BookingFields, managerId int64) (float32, error) {
	res, err := bs.bookingRepository.GetPredictions([]ent.BookingFields{bss}, managerId)
	if err != nil {
		return 0, err
	}
	return res[0], nil
}

func (bs BookingService) dateToString(b ent.Booking) string {
	return strings.Join([]string{strconv.FormatInt(b.ArrivalDateDayOfMonth, 10), b.ArrivalDateMonth, strconv.Itoa(int(b.ArrivalDateYear))}, " ")
}

func (bs BookingService) getDate(b ent.Booking) (time.Time, error) {
	date, err := time.Parse("2 January 2006", bs.dateToString(b))
	return date, err
}

func (bs BookingService) getDateByString(b string) time.Time {
	date, _ := time.Parse("2 January 2006", b)
	return date
}

func (bs BookingService) TrainModel(bookings []ent.BookingFields, ys []int64, managerId int64) error {
	fmt.Println("service train")
	m, err := bs.userRepository.GetManagerById(managerId)
	if err != nil {
		return err
	}
	return bs.bookingRepository.TrainModel(bookings, ys, m.HotelId)
}

// func (bs BookingService) GetPrevFutureBookings(managerId int) (prevBookings []ent.Booking, futureBookings []ent.Booking, err error) {
// 	bookings, err := bs.bookingRepository.GetAllForManager(managerId)
// 	if err != nil {
// 		fmt.Println(err)
// 		return make([]ent.Booking, 0), make([]ent.Booking, 0), err
// 	}
// 	arrivalDates := (funk.Map(bookings, func(b ent.Booking) time.Time {
// 		date, _ := bs.getDate(b)
// 		return date
// 	})).([]time.Time)
// 	for i, b := range bookings {
// 		arrivalDate := arrivalDates[i]
// 		isAfter := arrivalDate.Compare(time.Now())
// 		if isAfter == -1 || isAfter == 0 {
// 			prevBookings = append(prevBookings, b)
// 		} else {
// 			futureBookings = append(futureBookings, b)
// 		}
// 	}
// 	if err != nil {
// 		return make([]ent.Booking, 0), make([]ent.Booking, 0), err
// 	}
// 	return prevBookings, futureBookings, nil
// }

// func (bs BookingService) GetProfitStatistic(managerId int) (ProfitStatistic, error) {
// prevBookings, futureBookings, err := bs.GetPrevFutureBookings(managerId)
// if err != nil {
// 	return ProfitStatistic{}, err
// }
// prevProfitMap := make(map[string]float32)
// for _, b := range prevBookings {
// 	if !b.IsCanceled {
// 		prevProfitMap[bs.dateToString(b)] = prevProfitMap[bs.dateToString(b)] + 1
// 	}
// }
// v := make([]Profit, 0, len(prevProfitMap))

// for key, value := range prevProfitMap {
// 	v = append(v, Profit{
// 		Date:  key,
// 		Value: value,
// 	})
// }
// sort.Slice(v, func(i, j int) bool {
// 	date1 := bs.getDateByString(v[i].Date)
// 	date2 := bs.getDateByString(v[j].Date)
// 	return date1.Before(date2)
// })
// // mapstructure.Decode(prevProfitMap, &prevProfit)

// predictions, _ := bs.bookingRepository.GetPredictions(
// 	funk.Map(futureBookings, func(fb ent.Booking) ent.BookingFields {
// 		return fb.BookingFields
// 	}).([]ent.BookingFields))

// FutureCancProfitMap := make(map[string]float32)
// FutureNotCancProfitMap := make(map[string]float32)
// for i, b := range futureBookings {
// 	if predictions[i] >= 0.5 {
// 		FutureCancProfitMap[bs.dateToString(b)] = FutureCancProfitMap[bs.dateToString(b)] + 1
// 	} else {
// 		FutureNotCancProfitMap[bs.dateToString(b)] = FutureNotCancProfitMap[bs.dateToString(b)] + 1
// 	}
// }
// arr_canc := make([]Profit, 0, len(FutureCancProfitMap))
// arr_not_canc := make([]Profit, 0, len(FutureNotCancProfitMap))

// for key, value := range FutureCancProfitMap {
// 	arr_canc = append(arr_canc, Profit{
// 		Date:  key,
// 		Value: value,
// 	})
// }

// for key, value := range FutureNotCancProfitMap {
// 	arr_not_canc = append(arr_not_canc, Profit{
// 		Date:  key,
// 		Value: value,
// 	})
// }

// sort.Slice(arr_canc, func(i, j int) bool {
// 	date1 := bs.getDateByString(arr_canc[i].Date)
// 	date2 := bs.getDateByString(arr_canc[j].Date)
// 	return date1.Before(date2)
// })

// sort.Slice(arr_not_canc, func(i, j int) bool {
// 	date1 := bs.getDateByString(arr_not_canc[i].Date)
// 	date2 := bs.getDateByString(arr_not_canc[j].Date)
// 	return date1.Before(date2)
// })

// return ProfitStatistic{
// 	Prev: PreviousProfitStatistic{
// 		Profit: v,
// 	},
// 	Future: FutureProfitStatistic{
// 		NotCanceledProfit: arr_not_canc,
// 		CanceledProfit:    arr_canc,
// 	},
// }, nil
// }

type ProfitStatistic struct {
	Prev   PreviousProfitStatistic
	Future FutureProfitStatistic
}

type PreviousProfitStatistic struct {
	Profit []Profit
}

type FutureProfitStatistic struct {
	NotCanceledProfit []Profit
	CanceledProfit    []Profit
}

type Profit struct {
	Date  string
	Value float32
}
