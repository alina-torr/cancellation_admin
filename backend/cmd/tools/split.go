package main

import (
	ent "booking/entities"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

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
	// BookingId int64  `csv:"booking_id"`
	Hotel string `csv:"hotel"`
	ent.TrainBooking
}

func splitHotels() {
	bookings := ReadCsvFile("../../../hotel_booking.csv")
	resortBookings := make([]*CSV_SOURCE, 0)
	cityBookings := make([]*CSV_SOURCE, 0)
	id := int64(1)
	for _, b := range bookings {
		b.BookingId = id
		rand.Seed(time.Now().Unix())
		names := []string{
			"Михаил",
			"Кирилл",
			"Василиса",
			"Анатолий",
			"Екатерина",
		}
		numbers := []string{
			"+7(922)111-05-00",
			"+7(351)240-04-40",
			"+7(495)755-69-83",
			"+7(904)135-23-77",
			"+7(951)705-30-03",
		}
		name := names[rand.Intn(len(names))]
		number := numbers[rand.Intn(len(numbers))]
		b.AddInfo = fmt.Sprintf("%s, %s", name, number)
		id++
		b.ArrivalDateYear = b.ArrivalDateYear + 8
		if b.Hotel == "Resort Hotel" {
			resortBookings = append(resortBookings, b)
		}
		if b.Hotel == "City Hotel" {
			cityBookings = append(cityBookings, b)
		}
	}

	fileResort, err := os.Create("resort-hotel-add-info.csv")
	if err != nil {
		panic(err)
	}
	fileCity, err := os.Create("city-hotel-add-info.csv")
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
