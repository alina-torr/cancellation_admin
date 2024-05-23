package services

// import (
// 	"errors"
// 	"github.com/golang-jwt/jwt/v5"
// 	"lexique/internal/entities"
// 	"time"
// )

// type UserRepository interface {
// 	Create(ur entities.UserRegister) error
// 	GetUserByEmail(email string) (entities.User, error)
// 	GetUserById(id int) (entities.User, error)
// }

// type AuthService struct {
// 	repository UserRepository
// 	authConfig AuthConfig
// }

// func NewAuthService(repository UserRepository, authConfig AuthConfig) *AuthService {
// 	return &AuthService{
// 		repository: repository,
// 		authConfig: authConfig,
// 	}
// }

// func (as AuthService) Register(ur entities.UserRegister) error {
// 	// todo: hash password
// 	// todo: check if user with this email doesn't exists
// 	// todo: email confirm
// 	return as.repository.Create(ur)
// }

// func (as AuthService) Login(ul entities.UserLogin) (entities.LoginResponse, error) {
// 	user, err := as.repository.GetUserByEmail(ul.Email)
// 	if err != nil {
// 		return entities.LoginResponse{}, err
// 	}
// 	if ul.Password != user.Password {
// 		return entities.LoginResponse{}, errors.New("wrong email or password")
// 	}
// 	jwt, err := as.GetJWTtoken(user.Id)
// 	if err != nil {
// 		return entities.LoginResponse{}, errors.New("login error")
// 	}
// 	return entities.LoginResponse{AccessToken: jwt}, nil
// }

// func (as AuthService) GetJWTtoken(userId int) (string, error) {
// 	now := time.Now()
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: &jwt.NumericDate{Time: now.Add(time.Hour * 24)},
// 		},
// 		UserId: userId,
// 	})
// 	tokenString, err := token.SignedString([]byte(as.authConfig.JwtKey))
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, err
// }

// type JWTClaim struct {
// 	UserId int `json:"user_id"`
// 	jwt.RegisteredClaims
// }

// func (as AuthService) ValidateToken(signedToken string) (userId int, err error) {
// 	var jwtClaim JWTClaim
// 	token, err := jwt.ParseWithClaims(
// 		signedToken,
// 		&jwtClaim,
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(as.authConfig.JwtKey), nil
// 		},
// 	)
// 	if err != nil {
// 		return
// 	}
// 	if !token.Valid {
// 		err = errors.New("invalid token")
// 	}

// 	if jwtClaim.ExpiresAt.Unix() < time.Now().Unix() {
// 		err = errors.New("token expired")
// 		return
// 	}

// 	userId = jwtClaim.UserId
// 	return
// }

// func (as AuthService) GetUserById(id int) (user entities.User, err error) {
// 	return as.repository.GetUserById(id)
// }
