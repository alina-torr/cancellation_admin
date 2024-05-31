package services

import (
	ent "booking/entities"
	mlgrpc "booking/ml/grpc"
	"fmt"
	"github.com/thoas/go-funk"
	"os"
	"strconv"
	"strings"
	"time"
)

type bookingRepository interface {
	GetAllForManager(managerId int64) ([]ent.BookingTable, error)
	SaveBookingPredictions(booking []ent.BookingFields, predictions []float32, hotelId int64) error
	GetPredictions(bs []ent.BookingFields, hotelId int64) ([]float32, error)
	TrainModel(bookings []ent.BookingFields, cancellations []int64, hotelId int64) error
	GetApiKey(hotelId int64) string
	GetUserByApiKey(apiKey string) (int64, error)
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

func (bs BookingService) GetPredicts(bookings []*mlgrpc.Booking) {
	res := mlgrpc.BookingPredictRequest{
		Bookings: bookings,
	}
	fmt.Println(res)
}

func (bs BookingService) IsThereModel(managerId int64) bool {
	m, err := bs.userRepository.GetManagerById(managerId)
	if err != nil {
		return false
	}
	if _, err := os.Stat(fmt.Sprintf("../../files/model_%d.joblib", m.HotelId)); err == nil {
		return true
	}
	return false
}

func (bs BookingService) SaveBookingPredictions(booking []ent.BookingFields, predictions []float32, managerId int64) error {
	m, err := bs.userRepository.GetManagerById(managerId)
	if err != nil {
		return err
	}
	return bs.bookingRepository.SaveBookingPredictions(booking, predictions, m.HotelId)
}

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

func (bs BookingService) dateToString(b ent.BookingTable) string {
	return strings.Join([]string{strconv.FormatInt(b.ArrivalDateDayOfMonth, 10), b.ArrivalDateMonth, strconv.Itoa(int(b.ArrivalDateYear))}, " ")
}

func (bs BookingService) getDate(b ent.BookingTable) (time.Time, error) {
	date, err := time.Parse("2 January 2006", bs.dateToString(b))
	return date, err
}

func (bs BookingService) getDateByString(b string) time.Time {
	date, _ := time.Parse("2 January 2006", b)
	return date
}

func (bs BookingService) TrainModel(bookings []ent.BookingFields, ys []int64, managerId int64) error {
	m, err := bs.userRepository.GetManagerById(managerId)
	if err != nil {
		return err
	}
	return bs.bookingRepository.TrainModel(bookings, ys, m.HotelId)
}

func (bs BookingService) GetApiKey(managerId int64) string {
	m, _ := bs.userRepository.GetManagerById(managerId)
	return bs.bookingRepository.GetApiKey(m.HotelId)
}

func (bs BookingService) GetUserByApiKey(apiKey string) (int64, error) {
	return bs.bookingRepository.GetUserByApiKey(apiKey)
}

func (bs BookingService) GetFutureBookings(managerId int64) (futureBookings []ent.BookingTable, err error) {
	bookings, err := bs.bookingRepository.GetAllForManager(managerId)
	if err != nil {
		fmt.Println(err)
		return make([]ent.BookingTable, 0), err
	}
	arrivalDates := (funk.Map(bookings, func(b ent.BookingTable) time.Time {
		date, _ := bs.getDate(b)
		return date
	})).([]time.Time)
	for i, b := range bookings {
		arrivalDate := arrivalDates[i]
		isAfter := arrivalDate.Compare(time.Now())
		if isAfter != -1 && isAfter != 0 {
			futureBookings = append(futureBookings, b)
		}
	}
	if err != nil {
		return make([]ent.BookingTable, 0), err
	}
	return futureBookings, nil
}
