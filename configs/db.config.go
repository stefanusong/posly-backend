package configs

import (
	"fmt"

	"github.com/stefanusong/posly-backend/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDBConnection() *gorm.DB {
	dsn := utils.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.LogFatalIfError("Error opening db connection : ", err)
	fmt.Println("Connected to database")

	return db
}
