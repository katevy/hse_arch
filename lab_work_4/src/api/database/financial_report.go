package database

import "api/models"

type FinancialReport interface {
	CreateFinancialReport(history *models.FinancialReport) (*models.FinancialReport, error)
}
