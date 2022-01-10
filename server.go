package main

import (
	"github.com/stefanusong/posly-backend/configs"
	"github.com/stefanusong/posly-backend/helpers"
	"github.com/stefanusong/posly-backend/routes"
)

func main() {
	helpers.LoadEnv()
	e, db := routes.Init()
	defer configs.DisconnectDB(db)
	e.Logger.Fatal(e.Start(":8080"))
}
