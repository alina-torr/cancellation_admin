package entities

type Hotel struct {
	Id int `db:"id"`
	HotelFields
}

type Room struct {
	Id int `db:"id"`
	RoomFields
}

type HotelFields struct {
	Name    string `db:"name"`
	Country string `db:"country"`
	City    string `db:"city"`
}

type RoomFields struct {
	Hotel_id int     `db:"hotel_id"`
	Price    float32 `db:"price"`
}
