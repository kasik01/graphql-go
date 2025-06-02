package database

import (
	"fmt"
	"graphql-hasura-demo/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Db  *gorm.DB
	Gin *gin.Engine
}

func GetConfig() *Config {
	Connect()
	InitModels()

	return &Config{
		Db:  GetDB(),
		Gin: gin.Default(),
	}
}

var db *gorm.DB

func InitDb() {
	Connect()
	InitModels()
}

func Connect() {
	var appConfig = config.GetAppConfig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		appConfig.DB_HOST, appConfig.DB_USER, appConfig.DB_PASSWORD, appConfig.DB_NAME, appConfig.DB_PORT,
	)
	// dsn := "host=postgres user=postgres password=Khoa2401_ dbname=std_mng port=5432 sslmode=disable TimeZone=UTC"

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
