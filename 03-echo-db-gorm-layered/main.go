package main

import (
	"sirclo/restapi/db/gorm/datastore"
	"sirclo/restapi/db/gorm/delivery"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := datastore.InitDB()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	delivery.InitUserRoute(e, db)
	// routing with query parameter

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
