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

	r.POST("/login", userHandler.LoginHandler())
	r.POST("/register", userHandler.RegisterHandler())

	api := r.Group("/api")
	api.Use(middleware.TokenAuthMiddleware(*userService))
	api.GET("/is_there_model", bookingHandler.IsThereModel())
	api.POST("/upload_data", bookingHandler.UploadBookingFile())
	api.POST("/upload_data_predictions", bookingHandler.UploadBookingPredictionFile())
	api.GET("/get_predictions", bookingHandler.GetPredictionsBooking())
	api.POST("/prediction", bookingHandler.GetPredict())

	runSQL("../db/init.sql", dbpool)

	initTestData(hotelService, userService)

	r.Run(":5001")
}

func initTestData(hotelService *services.HotelService, userService *services.UserService) {

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
}

func connectDB() *pgxpool.Pool {
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
