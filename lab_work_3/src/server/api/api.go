package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"../service"
)

// предоставляет методы для работы с компаниями через API
type CompanyAPI struct {
	companyService service.CompanyService
}

// создает новый экземпляр CompanyAPI
func NewCompanyAPI(companyService service.CompanyService) *CompanyAPI {
	return &CompanyAPI{companyService}
}

//обрабатывает запрос на добавление компании в избранное
func (c *CompanyAPI) AddFavoriteCompanyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyName := vars["companyName"]

	err := c.companyService.AddFavoriteCompany(companyName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}