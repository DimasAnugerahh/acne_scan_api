package database

import (
	"acne-scan-api/configs"
	"acne-scan-api/internal/model/domain"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(config *configs.MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	migration(db)

	return db, nil
}

func migration(db *gorm.DB){
	db.AutoMigrate(
		&domain.Users{},
	)
}