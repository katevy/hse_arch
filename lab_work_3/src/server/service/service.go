package service

import "./database"

// предоставляет методы для работы с компаниями
type CompanyService interface {
	AddFavoriteCompany(companyName string) error
}

// реализует CompanyService
type CompanyServiceImpl struct {
	dbHandler database.DatabaseHandler
}

// создает новый экземпляр CompanyServiceImpl
func NewCompanyService(dbHandler database.DatabaseHandler) *CompanyServiceImpl {
	return &CompanyServiceImpl{dbHandler}
}

// добавляет компанию в избранное
func (s *CompanyServiceImpl) AddFavoriteCompany(companyName string) error {
	return s.dbHandler.AddFavoriteCompany(companyName)
}