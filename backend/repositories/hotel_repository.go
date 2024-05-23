package repositories

import (
	ent "booking/entities"
	"context"

	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HotelRepository struct {
	dbpool *pgxpool.Pool
}

func NewHotelRepository(dbpool *pgxpool.Pool) *HotelRepository {
	return &HotelRepository{
		dbpool: dbpool,
	}
}

func (hr HotelRepository) GetById(id int64, managerId int64) (ent.Hotel, error) {
	return ent.Hotel{}, nil
}

func (hr HotelRepository) Create(hotel ent.HotelFields) (int64, error) {
	tx, err := hr.dbpool.Begin(context.Background())
	if err != nil {
		return -1, err
	}
	defer tx.Rollback(context.Background())

	var hotelId int64
	err = tx.QueryRow(context.Background(),
		`INSERT INTO hotelscheme.hotel (name, country, city)
			VALUES ($1, $2, $3) RETURNING id`, hotel.Name, hotel.Country, hotel.City).Scan(&hotelId)

	if err != nil {
		return -1, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return -1, err
	}
	return hotelId, nil
}

func (hr HotelRepository) CreateRoom(room ent.RoomFields) (int64, error) {
	tx, err := hr.dbpool.Begin(context.Background())
	if err != nil {
		return -1, err
	}
	defer tx.Rollback(context.Background())

	var roomId int64
	err = tx.QueryRow(context.Background(),
		`INSERT INTO hotelscheme.room (hotel_id, price)
			VALUES ($1, $2) RETURNING id`, room.Hotel_id, room.Price).Scan(&roomId)

	if err != nil {
		return -1, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return -1, err
	}
	return roomId, nil
}
