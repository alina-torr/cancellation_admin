package main

import (
	ent "booking/entities"
	handlers "booking/http_handlers"
	"booking/middleware"
	"booking/repositories"
	"booking/services"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func String(v string) *string { return &v }

func main() {
	dbpool := connectDB()
	defer dbpool.Close()

	hotelRepository := repositories.NewHotelRepository(dbpool)
	hotelService := services.NewHotelService(hotelRepository)
	userRepository := repositories.NewUserRepository(dbpool)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewGinUserHandler(userService, hotelService)
	bookingRepository := repositories.NewBookingRepository(dbpool)
	bookingService := services.NewBookingService(bookingRepository, userRepository)
	bookingHandler := handlers.NewGinBookingHandler(bookingService)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           300,
	}))

	// todo: add refresh jwt token
	r.POST("/login", userHandler.LoginHandler())
	r.POST("/register", userHandler.RegisterHandler())

	api := r.Group("/api")
	api.Use(middleware.TokenAuthMiddleware(*userService))
	api.GET("/is_there_model", bookingHandler.IsThereModel())
	api.POST("/upload_data", bookingHandler.UploadBookingFile())
	api.POST("/upload_data_predictions", bookingHandler.UploadBookingPredictionFile())
	api.GET("/get_predictions", bookingHandler.GetPredictionsBooking())
	// todo: дообучить модель
	// api.GET("/bookings", bookingHandler.GetAllBookings())
	// api.GET("/statistics", bookingHandler.GetStatictic())
	// api.POST("/predictions", bookingHandler.GetPredicts())
	api.POST("/prediction", bookingHandler.GetPredict())

	runSQL("../db/init.sql", dbpool)

	initTestData(hotelService, userService, bookingService)

	r.Run(":5001")
}

func initTestData(hotelService *services.HotelService, userService *services.UserService, bookingService *services.BookingService) {

	idCityHotel, err := hotelService.Create(ent.HotelFields{
		Name:    "City Hotel",
		Country: "PRT",
		City:    "",
	})
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = userService.CreateManager(ent.ManagerData{
		Login:    "test2",
		Password: "123",
	}, idCityHotel)
	if err != nil {
		fmt.Println(err.Error())
	}

	// for _, b := range bookings {
	// 	var clientId int
	// 	clientByEmail, err1 := userService.GetClientByEmail(b.Email)
	// 	clientByPhone, err2 := userService.GetClientByPhone(b.PhoneNumber)
	// 	if err1 != nil && err2 != nil {
	// 		name := strings.Fields(b.Name)
	// 		clientId, err = userService.CreateClient(ent.ClientRegister{
	// 			Email:        b.Email,
	// 			Phone_number: b.PhoneNumber,
	// 			First_name:   name[0],
	// 			Last_name:    name[1],
	// 			Country:      b.Country,
	// 		})
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	} else if err1 == nil {
	// 		clientId = clientByEmail.Id
	// 	} else if err2 == nil {
	// 		clientId = clientByPhone.Id
	// 	}

	// 	var hotelId int
	// 	if b.Hotel == "Resort Hotel" {
	// 		hotelId = idResortHotel
	// 	} else {
	// 		hotelId = idCityHotel
	// 	}

	// 	idRoom, err := hotelService.CreateRoom(ent.RoomFields{
	// 		Hotel_id: hotelId,
	// 		Price:    b.Adr,
	// 	})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	var meal sql.NullString
	// 	if strings.ToLower(b.Meal) == "undefined" {
	// 		meal = sql.NullString{Valid: false}
	// 	} else {
	// 		meal = sql.NullString{String: b.Meal, Valid: true}
	// 	}
	// 	requestValues := []string{
	// 		"Lorem ipsum dolor sit amet, consectetur adipiscing elit",
	// 		"In at nisl ac sapien suscipit consectetur",
	// 		"Aenean venenatis ante ut pharetra fermentum",
	// 		"Aenean tincidunt, libero at sagittis lobortis, magna tortor facilisis est, et lacinia risus ante ac lorem",
	// 		"Morbi non dolor nisl",
	// 		"Donec elementum tincidunt risus, auctor volutpat tortor interdum et",
	// 		"Aliquam erat volutpat",
	// 		"Nullam sapien ex, tempor et nibh sit amet, malesuada sollicitudin orci",
	// 	}
	// 	specialRequests := funk.Map(make([]string, b.TotalOfSpecialRequests), func(x string) string {
	// 		return requestValues[rand.Intn(len(requestValues))]
	// 	}).([]string)

	// 	var agent sql.NullString
	// 	if b.Agent == 0 {
	// 		agent = sql.NullString{Valid: false}
	// 	} else {
	// 		agent = sql.NullString{String: strconv.Itoa(b.Agent), Valid: true}
	// 		// fmt.Println(agent)

	// 	}

	// 	var company sql.NullString
	// 	if b.Company == 0 {
	// 		company = sql.NullString{Valid: false}
	// 	} else {
	// 		company = sql.NullString{String: strconv.Itoa(b.Company), Valid: true}
	// 	}

	// 	var arrivalDate time.Time
	// 	day := strconv.Itoa(b.ArrivalDateDayOfMonth)
	// 	month := b.ArrivalDateMonth
	// 	year := strconv.Itoa(b.ArrivalDateYear - 2016 + time.Now().Year())
	// 	arrivalDate, err = time.Parse("2 January 2006", strings.Join([]string{day, month, year}, " "))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	bookingDate := arrivalDate.Add(time.Duration(b.Leadtime * -24 * (int(time.Hour))))
	// 	confirmedDate := bookingDate.Add(time.Duration(b.DaysInWaitingList * 24 * (int(time.Hour))))

	// 	_, err = bookingService.Create(ent.BookingFields{
	// 		ClientId:                 clientId,
	// 		HotelId:                  hotelId,
	// 		ReservedRoomId:           idRoom,
	// 		AssignedRoomId:           idRoom,
	// 		Adults:                   b.Adults,
	// 		Children:                 b.Children,
	// 		Babies:                   b.Babies,
	// 		NightStays:               b.StaysInWeekNights + b.StaysInWeekendNights,
	// 		SpecialRequests:          specialRequests,
	// 		Agent:                    agent,
	// 		Company:                  company,
	// 		RequiredCarParkingSpaces: b.RequiredCarParkingSpaces,
	// 		ArrivalDate:              arrivalDate,
	// 		BookingDate:              bookingDate,
	// 		ConfirmedDate:            confirmedDate,
	// 		IsCanceled:               b.IsCanceled,
	// 		Meal:                     meal,
	// 		MarketSegment:            b.MarketSegment,
	// 		DistributionChannel:      b.DistributionChannel,
	// 		CustomerType:             b.CustomerType,
	// 		Deposit:                  0,
	// 		BookingChanges:           b.BookingChanges,
	// 		ReservationStatus:        "created",
	// 	})
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	// fmt.Println(bookingId)
	// }
}

// todo: config
func connectDB() *pgxpool.Pool {
	// url := fmt.Sprintf("postgresql://postgres:@%s:%s/%s", config.DBConfig.User, config.DBConfig.Password,
	// 	config.DBConfig.Host, config.DBConfig.Port, config.DBConfig.Name)
	url := "postgresql://postgres:harr23el@localhost:5432/booking_hotel"
	dbpool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}

func runSQL(path string, dbpool *pgxpool.Pool) {
	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		fmt.Printf(ioErr.Error())
	}
	sql := string(c)
	_, err := dbpool.Exec(context.Background(), sql)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
