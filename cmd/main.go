package main

import (
	"acne-scan-api/configs"
	"acne-scan-api/internal/app"
	"acne-scan-api/internal/infrastructure/database"
	cloudstorage "acne-scan-api/internal/pkg/cloud_storage"
	"database/sql"
	"sync"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {

	var db *sql.DB
	// var config
	var err error
	var bucketUploader cloudstorage.StorageBucketUploader

	fiber := fiber.New()

	var wg sync.WaitGroup
	wg.Add(4)

	// go func() {
	// 	defer wg.Done()
	config, err := configs.LoadConfig()
	if err != nil {
		logrus.Fatal("Error loading config:", err.Error())
	}
	// }()

	// Initialize database connection
	go func() {
		defer wg.Done()
		db, err = database.NewMySQLConnection(&config.MySQL)
		if err != nil {
			logrus.Fatal("Error connecting to MySQL:", err.Error())
		}
	}()

	validate := validator.New()

	go func() {
		defer wg.Done()
		bucketUploader = cloudstorage.NewStorageBucketUploader(&config.StorageBucket)
	}()

	go func() {
		defer wg.Done()
		app.InitApp(db, validate, fiber, bucketUploader)
	}()

	fiber.Use(cors.New())
	fiber.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	go func() {
		defer wg.Done()
		fiber.Listen(":8080")
	}()
	
	wg.Wait()

}
