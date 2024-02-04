package sqlite

import "api/models"

func (db *SQLiteDB) CreateFinancialReport(report *models.FinancialReport) (*models.FinancialReport, error) {
	q := `INSERT INTO financial_report (title, content, date) VALUES (?, ?, ?)`

	result, err := db.DB.Exec(q, report.Title, report.Content, report.Date)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	report.ID = int(id)

	return report, nil
}
