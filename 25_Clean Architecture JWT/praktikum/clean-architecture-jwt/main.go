package main

import (
	"go-ca/config"
	"go-ca/route"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := route.New(db)
	app.Logger.Fatal(app.Start(":8080"))
}
