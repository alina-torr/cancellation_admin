package handlers

import (
	ent "booking/entities"
	"booking/functions"
	rep "booking/repositories"
	serv "booking/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/thoas/go-funk"
)

type bookingService interface {
	// GetAllForManager(managerId int64) ([]ent.Booking, error)
	// Create(booking ent.BookingFields) (int64, error)
	// CreateBookings(booking []ent.BookingFields) error
	// GetDistributionChannel(managerId int) ([]rep.CountStatictic, error)
	TrainModel(bookings []ent.BookingFields, cancellations []int64, managerId int64) error
	GetPredictions(bss []ent.BookingFields, managerId int64) ([]float32, error)
	GetPrediction(bss ent.BookingFields, managerId int64) (float32, error)
	// GetProfitStatistic(managerId int64) (serv.ProfitStatistic, error)
	// GetPrevFutureBookings(managerId int64) (prevBookings []ent.Booking, futureBookings []ent.Booking, err error)
	// CreateBookings(booking []ent.BookingFields) error
}

type GinBookingHandler struct {
	bookingService bookingService
}

func NewGinBookingHandler(service bookingService) *GinBookingHandler {
	return &GinBookingHandler{
		bookingService: service,
	}
}

// func (uh *GinBookingHandler) GetAllBookings() gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		_, futureBookings, err := uh.bookingService.GetPrevFutureBookings(functions.GetUserId(c))
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			handleError(c, http.StatusInternalServerError, err, false)
// 			return
// 		}
// 		predictions, err := uh.bookingService.GetPredictions(funk.Map(futureBookings, func(fb ent.Booking) ent.BookingFields {
// 			return fb.BookingFields
// 		}).([]ent.BookingFields))
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			handleError(c, http.StatusInternalServerError, err, false)
// 			return
// 		}
// 		res := make([]ent.BookingTable, 0)
// 		for i, b := range futureBookings {
// 			if predictions[i] > 0.5 {
// 				res = append(res, ent.CastBookingFieldsToTable(b.BookingFields, predictions[i]))
// 			}
// 		}
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			handleError(c, http.StatusInternalServerError, err, false)
// 			return
// 		}
// 		c.JSON(http.StatusOK, res)
// 	}
// }

func (uh *GinBookingHandler) UploadBookingPredictionFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		bookings := []*ent.TrainBooking{}

		f, _ := file.Open()
		if err := gocsv.UnmarshalMultipartFile(&f, &bookings); err != nil {
			fmt.Println(len(bookings))
			handleError(c, http.StatusInternalServerError, err, false)

		} else {
			fmt.Println(len(bookings))
			books := funk.Map(bookings, func(fb *ent.TrainBooking) ent.BookingFields {
				return fb.BookingFields
			}).([]ent.BookingFields)
			// cans := funk.Map(bookings, func(fb *ent.TrainBooking) int64 {
			// 	return fb.IsCanceled
			// }).([]int64)

			id := functions.GetUserId(c)
			res, err := uh.bookingService.GetPredictions(books, id)

			if err != nil {
				fmt.Println(err)
				handleError(c, http.StatusInternalServerError, err, false)
			}

			// _ = uh.bookingService.CreateBookings().([]ent.BookingFields))
			for _, b := range bookings {
				id, err := uh.bookingService.Create(ent.CastCSVtoDB(*b, 1)) // todo: add hotel id
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(id)
				}
			}

			c.JSON(http.StatusOK, res)
		}

	}
}

func (uh *GinBookingHandler) UploadBookingFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		bookings := []*ent.TrainBooking{}

		f, _ := file.Open()
		if err := gocsv.UnmarshalMultipartFile(&f, &bookings); err != nil {
			fmt.Println(len(bookings))
			handleError(c, http.StatusInternalServerError, err, false)

		} else {
			fmt.Println(len(bookings))
			books := funk.Map(bookings, func(fb *ent.TrainBooking) ent.BookingFields {
				return fb.BookingFields
			}).([]ent.BookingFields)
			cans := funk.Map(bookings, func(fb *ent.TrainBooking) int64 {
				return fb.IsCanceled
			}).([]int64)

			id := functions.GetUserId(c)
			fmt.Println(id)
			err = uh.bookingService.TrainModel(books, cans, id)
			fmt.Println(err)

			if err != nil {
				fmt.Println(err)
				handleError(c, http.StatusInternalServerError, err, false)
			}
			// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
			// _ = uh.bookingService.CreateBookings().([]ent.BookingFields))
			// for _, b := range bookings {
			// 	id, err := uh.bookingService.Create(ent.CastCSVtoDB(*b, 1)) // todo: add hotel id
			// 	if err != nil {
			// 		fmt.Println(err.Error())
			// 	} else {
			// 		fmt.Println(id)
			// 	}
			// }

		}

	}
}

// func (uh *GinBookingHandler) GetStatictic() gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		dc, err := uh.bookingService.GetDistributionChannel(functions.GetUserId(c))
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			handleError(c, http.StatusInternalServerError, err, false)
// 		}
// 		ps, err := uh.bookingService.GetProfitStatistic(functions.GetUserId(c))
// 		res := Statictic{
// 			DistributionChannel: dc,
// 			ProfitStat:          ps,
// 		}
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			handleError(c, http.StatusInternalServerError, err, false)
// 		} else {
// 			c.JSON(http.StatusOK, res)
// 		}
// 	}
// }

// func (uh *GinBookingHandler) GetPredicts() gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		var bookings []ent.BookingFields

// 		if err := c.BindJSON(&bookings); err != nil {
// 			handleError(c, http.StatusBadRequest, err, false)
// 			return
// 		}

// 		res, err := uh.bookingService.GetPredictions(bookings)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			handleError(c, http.StatusInternalServerError, err, false)
// 		} else {
// 			c.JSON(http.StatusOK, res)
// 		}
// 	}
// }

func (uh *GinBookingHandler) GetPredict() gin.HandlerFunc {

	return func(c *gin.Context) {
		var booking ent.BookingFields

		if err := c.BindJSON(&booking); err != nil {
			handleError(c, http.StatusBadRequest, err, false)
			return
		}
		id := functions.GetUserId(c)
		res, err := uh.bookingService.GetPrediction(booking, id)
		if err != nil {
			fmt.Println(err.Error())
			handleError(c, http.StatusInternalServerError, err, false)
		} else {
			c.JSON(http.StatusOK, res)
		}
	}
}

type Statictic struct {
	DistributionChannel []rep.CountStatictic
	ProfitStat          serv.ProfitStatistic
}
