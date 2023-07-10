package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/database"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/routes"
)

func main() {
	db := database.ConnectDB()
	if err := database.MigrateDB(db); err != nil {
		log.Fatal(err)
	}
	if err := database.SeedDB(db); err != nil {
		log.Fatal(err)
	}

	r := routes.Init(db)
	loggedRouter := gorillaHandlers.LoggingHandler(os.Stdout, r)

	srv := &http.Server{
		Handler:      gorillaHandlers.CompressHandler(loggedRouter),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("GO REST server running on 8080")

	log.Fatal(srv.ListenAndServe())
}
