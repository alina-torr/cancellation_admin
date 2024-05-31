package entities

type Manager struct {
	Id      int64 `db:"id"`
	HotelId int64 `db:"hotel_id"`
	ManagerData
}

type ManagerRegisterDB struct {
	HotelId int64 `db:"hotel_id"`
	ManagerRegister
}

type ManagerLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ManagerRegister struct {
	ManagerData
	HotelCity    string `json:"hotelCity" db:"city"`
	HotelCountry string `json:"hotelCountry" db:"country"`
	HotelName    string `json:"hotelName" db:"name"`
}

type ManagerData struct {
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type StartInfo struct {
	IsThereModel bool
	ApiKey       string
}
