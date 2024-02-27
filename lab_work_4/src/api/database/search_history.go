package database

import "api/models"

type SearchHistory interface {
	CreateSearchHistory(history *models.SearchHistory) (*models.SearchHistory, error)
	GetSearchHistoryByID(id int) (*models.SearchHistory, error)
	GetAllSeacrhHistory() ([]*models.SearchHistory, error)
	UpdateSearchHistory(history *models.SearchHistory) error
	DeleteSearchHistory(id int) error
}
