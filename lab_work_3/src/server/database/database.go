package database

import (
	"database/sql"
	"log"
)

//  предоставляет методы для работы с базой данных
type DatabaseHandler interface {
	AddFavoriteCompany(companyName string) error
}

// реализует DatabaseHandler для работы с SQLite
type SQLiteHandler struct {
	db *sql.DB
}

// создает новый экземпляр SQLiteHandler
func NewSQLiteHandler() *SQLiteHandler {
	db, err := sql.Open("sqlite3", "favorite_companies.db")
	if err != nil {
		log.Fatal(err)
	}
	return &SQLiteHandler{db}
}

// добавляет компанию в избранное
func (s *SQLiteHandler) AddFavoriteCompany(companyName string) error {
	_, err := s.db.Exec("INSERT INTO favorite_companies (name) VALUES (?)", companyName)
	return err
}
