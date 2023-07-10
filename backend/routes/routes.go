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

	r := mux.NewRouter()

	// APIs
	apis := r.PathPrefix("/api").Subrouter()
	apis.HandleFunc("/farms", farmHandler.GetAll).Methods("GET")
	apis.PathPrefix("/").Handler(http.NotFoundHandler())

	return r
}
