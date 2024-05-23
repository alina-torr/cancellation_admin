package main

import (
	ent "booking/entities"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func main() {
	splitHotels()
}

func ReadCsvFile(filePath string) []*CSV_SOURCE {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	bookings := []*CSV_SOURCE{}

	if err := gocsv.UnmarshalFile(f, &bookings); err != nil {
		panic(err)
	}

	return bookings
}

type CSV_SOURCE struct {
	BookingId int64  `csv:"booking_id"`
	Hotel     string `csv:"hotel"`
	ent.TrainBooking
}

func splitHotels() {
	bookings := ReadCsvFile("../../../hotel_booking.csv")
	resortBookings := make([]*CSV_SOURCE, 0)
	cityBookings := make([]*CSV_SOURCE, 0)
	id := int64(1)
	for _, b := range bookings {
		b.BookingId = id
		id++
		b.ArrivalDateYear = b.ArrivalDateYear + 8
		if b.Hotel == "Resort Hotel" {
			resortBookings = append(resortBookings, b)
		}
		if b.Hotel == "City Hotel" {
			cityBookings = append(cityBookings, b)
		}
	}

	fileResort, err := os.Create("resort-hotel.csv")
	if err != nil {
		panic(err)
	}
	fileCity, err := os.Create("city-hotel.csv")
	if err != nil {
		panic(err)
	}
	defer fileResort.Close()
	defer fileCity.Close()

	if err := gocsv.MarshalFile(&resortBookings, fileResort); err != nil {
		panic(err)
	}

	if err := gocsv.MarshalFile(&cityBookings, fileCity); err != nil {
		panic(err)
	}
}
