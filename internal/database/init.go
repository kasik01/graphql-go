package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	Connect()
	InitModels()
}

func Connect() {
	// var appConfig = config.GetAppConfig()
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
	// 	appConfig.DB_HOST, appConfig.DB_USER, appConfig.DB_PASSWORD, appConfig.DB_NAME, appConfig.DB_PORT,
	// )
	dsn := "host=localhost user=postgres password=Khoa2401_ dbname=std_mng port=5432 sslmode=disable TimeZone=UTC"

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	db = d
}

func GetDB() *gorm.DB {
	return db
}

// AutoMigrate
func InitModels() {
	db := GetDB()
	db.AutoMigrate(&Student{}, &Grade{})
}
