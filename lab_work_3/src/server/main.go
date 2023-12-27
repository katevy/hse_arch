package main

import (
	"log"
	"net/http"

	"./api"
	"./database"
	"./service"
)

func main() {
	dbHandler := database.NewDatabaseHandler()
	companyService := service.NewCompanyService(dbHandler)
	apiHandler := api.NewAPIHandler(companyService)

	router := mux.NewRouter()
	router.HandleFunc("/api/addFavoriteCompany/{companyName}", apiHandler.AddFavoriteCompanyHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}