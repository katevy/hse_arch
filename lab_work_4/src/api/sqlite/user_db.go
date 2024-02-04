package sqlite

import (
	"api/models"

	_ "github.com/mattn/go-sqlite3"
)

func (db *SQLiteDB) CreateUser(user *models.User) error {

	q := `INSERT INTO users (username, email) VALUES (?, ?)`

	_, err := db.DB.Exec(q, user.Username, user.Email)

	return err
}
