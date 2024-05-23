package services

// import (
// 	"fmt"
// 	"github.com/caarlos0/env"
// 	"github.com/joho/godotenv"
// )

// type DBConfig struct {
// 	Host     string `env:"LEXIQUE_DB_HOST,required"`
// 	Port     string `env:"LEXIQUE_DB_PORT,required"`
// 	User     string `env:"LEXIQUE_DB_USER,required"`
// 	Password string `env:"LEXIQUE_DB_PASSWORD,required"`
// 	Name     string `env:"LEXIQUE_DB_NAME,required"`
// }

// type AuthConfig struct {
// 	JwtKey string `env:"LEXIQUE_JWT_KEY,required"`
// }

// type Config struct {
// 	DBConfig
// 	AuthConfig
// }

// func NewConfig(envfile string) (congif *Config, err error) {
// 	congif = &Config{}
// 	err = godotenv.Load(envfile)

// 	if err != nil {
// 		fmt.Printf("Error loading %s file", envfile)
// 		return
// 	}

// 	dbCfg := DBConfig{}
// 	if err = env.Parse(&dbCfg); err != nil {
// 		fmt.Printf("Error parsing environment variables %+v\n", err)
// 		return
// 	}

// 	authCfg := AuthConfig{}
// 	if err = env.Parse(&authCfg); err != nil {
// 		fmt.Printf("Error parsing environment variables %+v\n", err)
// 		return
// 	}

// 	congif = &Config{
// 		DBConfig:   dbCfg,
// 		AuthConfig: authCfg,
// 	}
// 	return
// }
