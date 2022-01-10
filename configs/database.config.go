package configs

import (
	"os"

	"github.com/stefanusong/posly-backend/entities"
	"github.com/stefanusong/posly-backend/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// define data source name
func GetDSN() string {
	helpers.LoadEnv()
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	DSN := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPass

	return DSN
}

// open database connection
func ConnectDB() *gorm.DB {
	dsn := GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.HandleError(err)

	//Auto migration
	db.AutoMigrate(&entities.Resto{})

	return db
}

// clsoe database connection
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	helpers.HandleError(err)
	dbSQL.Close()
}
