package services

import (
	// "errors"
	// "github.com/golang-jwt/jwt/v5"
	"booking/consts"
	ent "booking/entities"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	// "time"
)

type userRepository interface {
	CreateManager(user ent.ManagerData, hotelId int64) (int64, error)
	CreateClient(user ent.ClientRegister) (int64, error)
	GetClientByPhone(phone_number string) (user ent.Client, err error)
	GetClientByEmail(email string) (user ent.Client, err error)
	GetManagerByLogin(login string) (user ent.Manager, err error)
	GetManagerById(id int64) (user ent.Manager, err error)
}

type UserService struct {
	repository userRepository
	// authConfig AuthConfig
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{
		repository: repository,
		// authConfig: authConfig,
	}
}

func (us UserService) CreateManager(user ent.ManagerData, hotelId int64) (int64, error) {
	return us.repository.CreateManager(user, hotelId)
}

func (us UserService) CreateClient(user ent.ClientRegister) (int64, error) {
	return us.repository.CreateClient(user)
}
func (us UserService) GetClientByPhone(phone_number string) (user ent.Client, err error) {
	return us.repository.GetClientByPhone(phone_number)
}
func (us UserService) GetClientByEmail(email string) (user ent.Client, err error) {
	return us.repository.GetClientByEmail(email)
}
func (us UserService) GetManagerByLogin(login string) (user ent.Manager, err error) {
	return us.repository.GetManagerByLogin(login)
}
func (us UserService) GetManagerById(id int64) (user ent.Manager, err error) {
	return us.repository.GetManagerById(id)
}

func (as UserService) Login(ul ent.ManagerLogin) (ent.LoginResponse, error) {
	user, err := as.repository.GetManagerByLogin(ul.Login)
	if err != nil {
		return ent.LoginResponse{}, errors.New("2")
	}
	if ul.Password != user.Password {
		return ent.LoginResponse{}, errors.New("2")
	}
	jwt, err := as.GetJWTtoken(user.Id)
	if err != nil {
		return ent.LoginResponse{}, errors.New("login error")
	}
	return ent.LoginResponse{AccessToken: jwt}, nil
}

// func (as AuthService) Register(ur entities.UserRegister) error {
// 	// todo: hash password
// 	// todo: check if user with this email doesn't exists
// 	// todo: email confirm
// 	return as.repository.Create(ur)
// }

func (as UserService) GetJWTtoken(userId int64) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: now.Add(time.Hour * 24)},
		},
		UserId: userId,
	})
	// todo: config JWT
	tokenString, err := token.SignedString([]byte(consts.JWT_PRIVATE_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

type JWTClaim struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func (us UserService) ValidateToken(signedToken string) (userId int64, err error) {
	var jwtClaim JWTClaim
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaim,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(consts.JWT_PRIVATE_KEY), nil
		},
	)
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("invalid token")
	}

	if jwtClaim.ExpiresAt.Unix() < time.Now().Unix() {
		err = errors.New("token expired")
		return
	}

	userId = jwtClaim.UserId
	return
}
