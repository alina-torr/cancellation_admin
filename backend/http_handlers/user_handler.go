package handlers

import (
	ent "booking/entities"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userService interface {
	Login(ul ent.ManagerLogin) (ent.LoginResponse, error)
	CreateManager(user ent.ManagerData, hotelId int64) (int64, error)
	GetManagerByLogin(login string) (ent.Manager, error)
}

type hotelService interface {
	Create(w ent.HotelFields) (int64, error)
}

type GinUserHandler struct {
	userService  userService
	hotelService hotelService
}

func NewGinUserHandler(service userService, hotelService hotelService) *GinUserHandler {
	return &GinUserHandler{
		userService:  service,
		hotelService: hotelService,
	}
}

func (h *GinUserHandler) LoginHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		var userData ent.ManagerLogin

		if err := c.BindJSON(&userData); err != nil {
			handleError(c, http.StatusBadRequest, err, false)
			return
		}

		lr, err := h.userService.Login(userData)
		if err != nil {
			handleError(c, http.StatusInternalServerError, err, true)
		} else {
			c.JSON(http.StatusOK, lr)
		}
	}
}

func (h *GinUserHandler) RegisterHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		var userData ent.ManagerRegister

		if err := c.BindJSON(&userData); err != nil {
			handleError(c, http.StatusBadRequest, err, false)
			return
		}

		manager, err := h.userService.GetManagerByLogin(userData.Login)
		if manager.Login != "" {
			handleError(c, http.StatusInternalServerError, errors.New("1"), true)
			return
		}
		hotelId, err := h.hotelService.Create(ent.HotelFields{
			Name:    userData.HotelName,
			Country: userData.HotelCountry,
			City:    userData.HotelCity,
		})
		if err != nil {
			fmt.Println(err.Error())
			handleError(c, http.StatusInternalServerError, err, true)
			return
		}
		lr, err := h.userService.CreateManager(ent.ManagerData{
			Login:    userData.Login,
			Password: userData.Password,
		}, hotelId)

		if err != nil {
			handleError(c, http.StatusInternalServerError, err, true)
			return
		} else {
			c.JSON(http.StatusOK, lr)
		}
	}
}
