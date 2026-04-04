package configs

import (
	"os"
	"fmt"
	"strconv"
)

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	Name string
	JWTExpirationInSeconds int64
	JWTSecret string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		JWTExpirationInSeconds: getEnvAsInt64("JWT_EXP", 3600*24*7),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}

func (c DBConfig) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name,
	)
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return defaultValue
	}

	return value
}