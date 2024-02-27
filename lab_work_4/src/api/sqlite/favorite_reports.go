package sqlite

import (
	"api/models"
	"errors"
)

func (db *SQLiteDB) CreateFavoriteReport(favorite *models.FavoriteReport) (*models.FavoriteReport, error) {
	q := `INSERT INTO favorite_report (user_id, report_id) VALUES (?, ?)`

	result, err := db.DB.Exec(q, 1, favorite.ReportID)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	favorite.ID = int(id)
	favorite.UserID = 1

	return favorite, nil
}

func (db *SQLiteDB) GetAllFavoriteReports() ([]*models.FavoriteReport, error) {
	var favoriteReports []*models.FavoriteReport

	q := `SELECT * FROM favorite_report WHERE user_id = ?`

	rows, err := db.DB.Query(q, 1)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var favorite models.FavoriteReport
		if err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.ReportID); err != nil {
			return nil, err
		}
		favoriteReports = append(favoriteReports, &favorite)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return favoriteReports, nil
}

func (db *SQLiteDB) GetFavoriteReportByID(id int) (*models.FavoriteReport, error) {
	var favoriteReport models.FavoriteReport

	q := `SELECT * FROM favorite_report WHERE user_id = ? AND id = ?  LIMIT 1`

	err := db.DB.QueryRow(q, 1, id).Scan(&favoriteReport.ID, &favoriteReport.UserID, &favoriteReport.ReportID)
	if err != nil {
		return nil, err
	}

	return &favoriteReport, nil
}

func (db *SQLiteDB) UpdateFavoriteReport(favorite *models.FavoriteReport) error {

	var count int
	checkQuery := `SELECT COUNT(*) FROM favorite_report WHERE user_id = ? AND id = ?`
	getErr := db.DB.Get(&count, checkQuery, 1, favorite.ID)
	if getErr != nil {
		return getErr
	}

	if count == 0 {
		return errors.New("record not found")
	}

	q := `UPDATE favorite_report SET report_id = ? WHERE user_id = ? AND id = ?`

	_, err := db.DB.Exec(q, favorite.ReportID, 1, favorite.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteDB) DeleteFavoriteReport(id int) error {

	var count int
	checkQuery := `SELECT COUNT(*) FROM favorite_report WHERE user_id = ? AND id = ?`
	err := db.DB.Get(&count, checkQuery, 1, id)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("record not found")
	}

	q := `DELETE FROM favorite_report WHERE user_id = ? AND id = ?`

	_, deleteError := db.DB.Exec(q, 1, id)
	if deleteError != nil {
		return deleteError
	}

	return nil
}
