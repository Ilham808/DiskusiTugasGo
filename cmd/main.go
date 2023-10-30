package main

import (
	route "DiskusiTugas/api/routes"
	"DiskusiTugas/config"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	app := config.InitConfig()
	db := config.InitDB(app)
	config.AutoMigrate(db)

	e := echo.New()
	route.SetupRoute(e, app, db)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8001)).Error())
}
