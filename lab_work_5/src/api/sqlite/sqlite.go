package sqlite

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	DB *sqlx.DB
}

func newSQLiteDB(db *sqlx.DB) *SQLiteDB {
	return &SQLiteDB{DB: db}
}

func (db *SQLiteDB) InitDB() error {
	//users
	if _, err := db.DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		email TEXT
	)`); err != nil {
		return err
	}

	//serchHistory
	if _, err := db.DB.Exec(`CREATE TABLE IF NOT EXISTS serch_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		content TEXT,
		date TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`); err != nil {
		return err
	}

	//financialReport
	if _, err := db.DB.Exec(`CREATE TABLE IF NOT EXISTS financial_report (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		content TEXT,
		date TIMESTAMP
	)`); err != nil {
		return err
	}

	//reportAnalysis
	if _, err := db.DB.Exec(`CREATE TABLE IF NOT EXISTS report_analysis (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		report_id INTEGER,
		result TEXT,
		FOREIGN KEY (report_id) REFERENCES financial_report(id)
	)`); err != nil {
		return err
	}

	//favoriteReport
	if _, err := db.DB.Exec(`CREATE TABLE IF NOT EXISTS favorite_report (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		report_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id),
		FOREIGN KEY (report_id) REFERENCES financial_report(id)
	)`); err != nil {
		return err
	}

	return nil
}

func ConnectDB(dbPath string) (*SQLiteDB, error) {
	db, err := sqlx.Connect("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}

	dbHandler := newSQLiteDB(db)

	if err := dbHandler.InitDB(); err != nil {
		log.Fatalln("Failed to initialize database:", err)
	}

	return dbHandler, nil
}
