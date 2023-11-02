package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DBUsername             string
	DBPassword             string
	DBPort                 string
	DBHost                 string
	DBName                 string
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	AccessTokenSecret      string
	RefreshTokenSecret     string
	CloudinaryAPISecret    string
	CloudinaryAPIKey       string
	CloudinaryCloudName    string
	CloudinaryUploadFolder string
	OpenAIKey              string
}

func InitConfig() *Config {
	var res = new(Config)
	res, err := LoadConfig()

	if err != nil {
		logrus.Error("Cannot load config, ", err.Error())
		return nil
	}

	return res
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	config := &Config{
		DBUsername:             os.Getenv("DB_USERNAME"),
		DBPassword:             os.Getenv("DB_PASSWORD"),
		DBPort:                 os.Getenv("DB_PORT"),
		DBHost:                 os.Getenv("DB_HOST"),
		DBName:                 os.Getenv("DB_NAME"),
		AccessTokenExpiryHour:  getEnvAsInt("ACCESS_TOKEN_EXPIRY_HOUR", 2),
		RefreshTokenExpiryHour: getEnvAsInt("REFRESH_TOKEN_EXPIRY_HOUR", 168),
		AccessTokenSecret:      os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:     os.Getenv("REFRESH_TOKEN_SECRET"),
		CloudinaryAPISecret:    os.Getenv("CLOUDINARY_API_SECRET"),
		CloudinaryAPIKey:       os.Getenv("CLOUDINARY_API_KEY"),
		CloudinaryCloudName:    os.Getenv("CLOUDINARY_CLOUD_NAME"),
		CloudinaryUploadFolder: os.Getenv("CLOUDINARY_UPLOAD_FOLDER"),
		OpenAIKey:              os.Getenv("OPEN_AI_KEY"),
	}

	return config, nil
}

func getEnvAsInt(key string, defaultValue int) int {
	if val, ok := os.LookupEnv(key); ok {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return defaultValue
}
