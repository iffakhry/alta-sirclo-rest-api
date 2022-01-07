package main

import (
	"fmt"
	"os"
	"sirclo/restapi/db/gorm/datastore"
	"sirclo/restapi/db/gorm/delivery"

	"github.com/labstack/echo/v4"
)

func main() {
	jwtSecret := os.Getenv("JWT_SECRET")
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	fmt.Println(connectionString)
	db, err := datastore.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	if err = datastore.InitialMigration(db); err != nil {
		panic(err)
	}
	e := echo.New()
	delivery.InitUserRoute(e, db, jwtSecret)
	// routing with query parameter

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
