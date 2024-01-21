package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	DbHost       string
	DbPort       string
	DbName       string
	DbUser       string
	DbPass       string
	DbSslMode    string
	AllowOrigins string
	AppPort      string
	Url          string
	ApiKey       string
	AgeUrl       string
	GenderUrl    string
	CountryUrl   string
}

var conf Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warning("Couldn't load env variables to connect to DB...")
	}

	conf = Config{
		AppPort:      getEnv("APP_PORT", ""),
		Url:          getEnv("URL", ""),
		ApiKey:       getEnv("API_KEY", ""),
		DbHost:       getEnv("DB_HOST", ""),
		DbPort:       getEnv("DB_PORT", ""),
		DbName:       getEnv("DB_NAME", ""),
		DbUser:       getEnv("DB_USER", ""),
		DbPass:       getEnv("DB_PASS", ""),
		DbSslMode:    getEnv("DB_SSLMODE", ""),
		AllowOrigins: getEnv("ALLOW_ORIGINS", ""),
		AgeUrl:       getEnv("AGE_URL", ""),
		GenderUrl:    getEnv("GENDER_URL", ""),
		CountryUrl:   getEnv("COUNTRY_URL", ""),
	}
	log.Info("configuration loaded")
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetConfigurationInstance() Config {
	return conf
}
