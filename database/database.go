package database

import (
	"fmt"
	"os"
	"time"

	"github.com/souzamoab/products-api-golang/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDatabase() {
	db_url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	//db_url := "host=postgres port=5432 user=sd dbname=productsdb sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
