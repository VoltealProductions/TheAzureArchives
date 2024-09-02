package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost  string
	Port        string
	DBDriver    string
	DBUser      string
	DBPassword  string
	DBAddress   string
	DBName      string
	JWTExpInSec int64
	JWTSecret   string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	return Config{
		PublicHost:  getEnv("PUBLIC_HOST", "localhost"),
		Port:        getEnv("PUBLIC_PORT", "3030"),
		DBDriver:    getEnv("DB_DRIVER", "sqlite3"),
		DBUser:      getEnv("DB_USER", "root"),
		DBPassword:  getEnv("DB_PASS", "password"),
		DBAddress:   fmt.Sprintf("%s:%s", getEnv("PUBLIC_HOST", "localhost"), getEnv("PUBLIC_PORT", "3030")),
		DBName:      getEnv("DB_NAME", "database"),
		JWTExpInSec: getEnvAsInt("JWT_EXPIRATION", 3600*24*7),
		JWTSecret:   getEnv("JWT_SECRET", "Secret-Goes-Here"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if val, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
