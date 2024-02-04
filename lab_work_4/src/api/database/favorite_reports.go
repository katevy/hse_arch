package database

import "api/models"

type FavoriteReport interface {
	CreateFavoriteReport(history *models.FavoriteReport) (*models.FavoriteReport, error)
	GetFavoriteReportByID(id int) (*models.FavoriteReport, error)
	GetAllFavoriteReports() ([]*models.FavoriteReport, error)
	UpdateFavoriteReport(history *models.FavoriteReport) error
	DeleteFavoriteReport(id int) error
}
