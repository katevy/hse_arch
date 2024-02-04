package sqlite

import (
	"api/models"
	"errors"
)

func (db *SQLiteDB) CreateSearchHistory(history *models.SearchHistory) (*models.SearchHistory, error) {
	q := `INSERT INTO serch_history (user_id, content, date) VALUES (?, ?, ?)`

	result, err := db.DB.Exec(q, 1, history.Content, history.Date)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	history.ID = int(id)
	history.UserID = 1

	return history, nil
}

func (db *SQLiteDB) GetAllSeacrhHistory() ([]*models.SearchHistory, error) {
	var searchHistories []*models.SearchHistory

	q := `SELECT * FROM serch_history WHERE user_id = ?`

	rows, err := db.DB.Query(q, 1)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var history models.SearchHistory
		if err := rows.Scan(&history.ID, &history.UserID, &history.Content, &history.Date); err != nil {
			return nil, err
		}
		searchHistories = append(searchHistories, &history)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return searchHistories, nil
}

func (db *SQLiteDB) GetSearchHistoryByID(id int) (*models.SearchHistory, error) {
	var searchHistory models.SearchHistory

	q := `SELECT * FROM serch_history WHERE user_id = ? AND id = ?  LIMIT 1`

	err := db.DB.QueryRow(q, 1, id).Scan(&searchHistory.ID, &searchHistory.UserID, &searchHistory.Content, &searchHistory.Date)
	if err != nil {
		return nil, err
	}

	return &searchHistory, nil
}

func (db *SQLiteDB) UpdateSearchHistory(history *models.SearchHistory) error {

	q := `UPDATE serch_history SET content = ?, date = ? WHERE user_id = ? AND id = ?`

	_, err := db.DB.Exec(q, history.Content, history.Date, 1, history.ID)
	if err != nil {
		return err
	}
	return nil
}

//TODO добавить проверку на наличие записи, возвращать ошибку если уже удалили

func (db *SQLiteDB) DeleteSearchHistory(id int) error {

	var count int
	checkQuery := `SELECT COUNT(*) FROM serch_history WHERE user_id = ? AND id = ?`
	err := db.DB.Get(&count, checkQuery, 1, id)
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("record not found")
	}

	q := `DELETE FROM serch_history WHERE user_id = ? AND id = ?`

	_, deleteError := db.DB.Exec(q, 1, id)
	if deleteError != nil {
		return deleteError
	}

	return nil
}
