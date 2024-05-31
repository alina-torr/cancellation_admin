package entities

import "time"

type Hotel struct {
	Id int `db:"id"`
	HotelFields
}

type HotelFields struct {
	Name    string `db:"name"`
	Country string `db:"country"`
	City    string `db:"city"`
}

type HotelInfo struct {
	ApiKey            string    `db:"api_key"`
	LastUpdatePredict time.Time `db:"last_update_predict"`
	LastUpdateTrain   time.Time `db:"last_update_train"`
}

type HotelDB struct {
	Id int `db:"id"`
	HotelFields
	HotelInfo
}
