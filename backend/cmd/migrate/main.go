package main

import (
	"log"

	"github.com/umitron-mission/sw-farm-full-stack-coding-template/database"
)

func main() {
	db := database.ConnectDB()
	// drop all
	if err := database.DropDB(db); err != nil {
		log.Fatal(err)
	}
	if err := database.MigrateDB(db); err != nil {
		log.Fatal(err)
	}
	if err := database.SeedDB(db); err != nil {
		log.Fatal(err)
	}
}
