package services

import (
	"booking/consts"
	ent "booking/entities"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type userRepository interface {
	CreateManager(user ent.ManagerData, hotelId int64) (int64, error)
	GetManagerByLogin(login string) (user ent.Manager, err error)
	GetManagerById(id int64) (user ent.Manager, err error)
}

type UserService struct {
	repository userRepository
}

func NewUserService(repository userRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (us UserService) CreateManager(user ent.ManagerData, hotelId int64) (int64, error) {
	return us.repository.CreateManager(user, hotelId)
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

func (as UserService) GetJWTtoken(userId int64) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: now.Add(time.Hour * 24)},
		},
		UserId: userId,
	})
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
