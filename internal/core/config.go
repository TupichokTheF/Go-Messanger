package core

import (
	"fmt"
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	BasePath string `env:"BASE_PATH" env-required:"true"`
	HTTPConfig
	DataBaseConfig
	LocalConfig
	JWTConfig
}

type HTTPConfig struct {
	HTTPHost string `env:"HTTP_HOST"`
	HTTPPort int    `env:"HTTP_PORT"`
}

func (conf *HTTPConfig) GetAddress() string {
	return fmt.Sprintf("%s:%v", conf.HTTPHost, conf.HTTPPort)
}

type DataBaseConfig struct {
	DBHost     string `env:"DB_HOST" env-required:"true"`
	DBPort     int    `env:"DB_PORT" env-required:"true"`
	DBUser     string `env:"DB_USER" env-required:"true"`
	DBPassword string `env:"DB_PASSWORD" env-required:"true"`
	DBName     string `env:"DB_NAME" env-required:"true"`
}

func (d *DataBaseConfig) GetURL() string {
	url := fmt.Sprintf("postgres://%s:%s@%s:%v/%s",
		d.DBUser, d.DBPassword, d.DBHost, d.DBPort, d.DBName)

	return url
}

type LocalConfig struct {
	Swagger bool `env:"SWAGGER" env-required:"false"`
}

type JWTConfig struct {
	SecretKet []byte `env:"JWT_SECRET_KEY" env-required:"true"`
	AccessTTL time.Duration `env:"ACCESS_TOKEN_TTL" env-required:"true"`
	RefreshTTL time.Duration `env:"REFRESH_TOKEN_TTL" env-required:"true"`
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Can't load env file")
	}

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal("Can't read config file")
	}

	return &cfg
}
