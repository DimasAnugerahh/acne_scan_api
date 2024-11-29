package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type appConfig struct {
	MySQL         MySQLConfig
	StorageBucket StorageBucketConfig
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type StorageBucketConfig struct {
	BucketName string
	GOOGLE_APPLICATION_CREDENTIALS string
}

func LoadConfig() (*appConfig, error) {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("failed to load environment variables from .env file: %w", err)
		}
		fmt.Println("db connection successfully")
	}

	return &appConfig{
		MySQL: MySQLConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Database: os.Getenv("DB_NAME"),
		},
		StorageBucket: StorageBucketConfig{
			BucketName: os.Getenv("BUCKET_NAME"),
			GOOGLE_APPLICATION_CREDENTIALS: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		},
	}, nil
}
