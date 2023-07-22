package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/handlers"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/repositories"
	"github.com/umitron-mission/sw-farm-full-stack-coding-template/services"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *mux.Router {
	farmRepo := repositories.NewFarmRepository(db)
	farmService := services.NewFarmService(farmRepo)
	farmHandler := handlers.NewFarmHandler(farmService)

	fishPenRepo := repositories.NewFishPenRepository(db)
	fishPenService := services.NewFishPenService(farmRepo, fishPenRepo)
	fishPenHandler := handlers.NewFishPenHandler(fishPenService)

	r := mux.NewRouter()

	// APIs
	apis := r.PathPrefix("/api").Subrouter()
	apis.HandleFunc("/farms", farmHandler.GetAll).Methods("GET")
	apis.HandleFunc("/farms/{farmID}", farmHandler.GetByID).Methods("GET")
	apis.HandleFunc("/farms/{farmID}/fishpens", fishPenHandler.Create).Methods("POST")
	apis.HandleFunc("/farms/{farmID}/fishpens", fishPenHandler.GetAllForFarm).Methods("GET")
	apis.HandleFunc("/farms/{farmID}/fishpens/{fishPenID}", fishPenHandler.GetByID).Methods("GET")
	apis.HandleFunc("/farms/{farmID}/fishpens/{fishPenID}", fishPenHandler.Update).Methods("PUT")
	apis.HandleFunc("/farms/{farmID}/fishpens/{fishPenID}", fishPenHandler.Delete).Methods("DELETE")
	apis.PathPrefix("/").Handler(http.NotFoundHandler())

	return r
}
