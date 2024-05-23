package entities

type Client struct {
	Id int64 `db:"id"`
	ClientRegister
}

type ClientRegister struct {
	Phone_number string `db:"phone_number"`
	Email        string `db:"email"`
	First_name   string `db:"first_name"`
	Last_name    string `db:"last_name"`
	Middle_name  string `db:"middle_name"`
	Country      string `db:"country"`
}

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
