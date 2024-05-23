package services

import (
	ent "booking/entities"
)

type hotelRepository interface {
	GetById(id int64, managerId int64) (ent.Hotel, error)
	Create(hotel ent.HotelFields) (int64, error)
	CreateRoom(room ent.RoomFields) (int64, error)
}

type HotelService struct {
	repository hotelRepository
}

func NewHotelService(repository hotelRepository) *HotelService {
	return &HotelService{
		repository: repository,
	}
}

func (ws HotelService) GetById(id int64, managerId int64) (ent.Hotel, error) {
	return ws.repository.GetById(id, managerId)
}

func (ws HotelService) Create(w ent.HotelFields) (int64, error) {
	return ws.repository.Create(w)
}

// func (ws HotelService) CreateRoom(room ent.RoomFields) (int, error) {
// 	return ws.repository.CreateRoom(room)
// }
