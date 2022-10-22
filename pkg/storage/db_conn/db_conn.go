package db_conn

import (
	"f_gin/pkg/service/config_loader"
	"f_gin/pkg/storage/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var host string
var user string
var password string
var dbName string
var port int
var sslmode string
var timeZone string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	host = os.Getenv("DB_HOST")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")

	ymlConfig := config_loader.GetDbYamlConfig()
	port = ymlConfig.Port
	sslmode = ymlConfig.SSLMode
	timeZone = ymlConfig.TimeZone
}

func migrateModels(dbConn *gorm.DB) {
	dbConn.AutoMigrate(&models.User{})
}

func GetDbConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbName,
		port,
		sslmode,
		timeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migrateModels(db)
	return db
}
