package repositories

import (
	ent "booking/entities"
	"context"
	"fmt"
	"math/rand"

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
	key, err := GenerateRandomString(48)
	var hotelId int64
	err = tx.QueryRow(context.Background(),
		`INSERT INTO hotelscheme.hotel (name, country, city, api_key)
			VALUES ($1, $2, $3, $4) RETURNING id`, hotel.Name, hotel.Country, hotel.City, key).Scan(&hotelId)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return -1, err
	}
	return hotelId, nil
}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num := rand.Intn(len(letters))
		ret[i] = letters[num]
	}

	return string(ret), nil
}
