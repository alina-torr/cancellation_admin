package handlers

import (
	ent "booking/entities"
	"booking/functions"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/thoas/go-funk"
)

type bookingService interface {
	SaveBookingPredictions(booking []ent.BookingFields, predictions []float32, managerId int64) error
	TrainModel(bookings []ent.BookingFields, cancellations []int64, managerId int64) error
	IsThereModel(managerId int64) bool
	GetApiKey(managerId int64) string
	GetPredictions(bss []ent.BookingFields, managerId int64) ([]float32, error)
	GetPrediction(bss ent.BookingFields, managerId int64) (float32, error)
	GetFutureBookings(managerId int64) (futureBookings []ent.BookingTable, err error)
	GetUserByApiKey(apiKey string) (int64, error)
}

type GinBookingHandler struct {
	bookingService bookingService
}

func NewGinBookingHandler(service bookingService) *GinBookingHandler {
	return &GinBookingHandler{
		bookingService: service,
	}
}

func (uh *GinBookingHandler) GetPredictionsBooking() gin.HandlerFunc {

	return func(c *gin.Context) {
		futureBookings, err := uh.bookingService.GetFutureBookings(functions.GetUserId(c))
		if err != nil {
			fmt.Println(err.Error())
			handleError(c, http.StatusInternalServerError, err, false)
			return
		}
		res := make([]ent.BookingTable, 0)
		for _, b := range futureBookings {
			if b.CancellationPredict > 0.5 {
				res = append(res, b)
			}
		}
		c.JSON(http.StatusOK, res)
	}
}

func (uh *GinBookingHandler) IsThereModel() gin.HandlerFunc {

	return func(c *gin.Context) {
		id := functions.GetUserId(c)
		isThereModel := uh.bookingService.IsThereModel(id)
		if isThereModel {
			c.JSON(http.StatusOK, &ent.StartInfo{
				IsThereModel: isThereModel,
				ApiKey:       uh.bookingService.GetApiKey(id),
			})
		} else {
			c.JSON(http.StatusOK, &ent.StartInfo{
				IsThereModel: isThereModel,
				ApiKey:       "",
			})
		}

	}
}

func (uh *GinBookingHandler) UploadBookingPredictionFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		bookings := []*ent.TrainBooking{}

		f, _ := file.Open()
		if err := gocsv.UnmarshalMultipartFile(&f, &bookings); err != nil {
			handleError(c, http.StatusInternalServerError, err, false)

		} else {
			books := funk.Map(bookings, func(fb *ent.TrainBooking) ent.BookingFields {
				return fb.BookingFields
			}).([]ent.BookingFields)

			id := functions.GetUserId(c)
			res, err := uh.bookingService.GetPredictions(books, id)
			print(res)
			if err != nil {
				fmt.Println(err)
				handleError(c, http.StatusInternalServerError, err, false)
			}

			err = uh.bookingService.SaveBookingPredictions(books, res, id)
			if err != nil {
				fmt.Println(err)
				handleError(c, http.StatusInternalServerError, err, false)
			}
			c.JSON(http.StatusOK, res)
		}

	}
}

func (uh *GinBookingHandler) UploadBookingFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		file, _ := c.FormFile("file")

		bookings := []*ent.TrainBooking{}

		f, _ := file.Open()
		if err := gocsv.UnmarshalMultipartFile(&f, &bookings); err != nil {
			handleError(c, http.StatusInternalServerError, err, false)

		} else {
			books := funk.Map(bookings, func(fb *ent.TrainBooking) ent.BookingFields {
				return fb.BookingFields
			}).([]ent.BookingFields)
			cans := funk.Map(bookings, func(fb *ent.TrainBooking) int64 {
				return fb.IsCanceled
			}).([]int64)

			id := functions.GetUserId(c)
			err = uh.bookingService.TrainModel(books, cans, id)

			if err != nil {
				fmt.Println(err)
				handleError(c, http.StatusInternalServerError, err, false)
			}
		}

	}
}

func (uh *GinBookingHandler) GetPredict() gin.HandlerFunc {

	return func(c *gin.Context) {
		var booking ent.BookingFields
		k := c.Query("key")
		if err := c.BindJSON(&booking); err != nil {
			handleError(c, http.StatusBadRequest, err, false)
			return
		}
		id, err := uh.bookingService.GetUserByApiKey(k)
		if err != nil {
			handleError(c, http.StatusInternalServerError, err, false)
		}
		res, err := uh.bookingService.GetPrediction(booking, id)
		if err != nil {
			fmt.Println(err.Error())
			handleError(c, http.StatusInternalServerError, err, false)
		} else {
			c.JSON(http.StatusOK, res)
		}
	}
}
