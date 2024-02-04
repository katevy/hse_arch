package models

import (
	"fmt"
	"time"
)

type User struct {
	ID       int
	Username string
	Email    string
}

type SearchHistory struct {
	ID      int
	UserID  int
	Content string
	Date    time.Time
}

func (s SearchHistory) String() string {
	return fmt.Sprintf("ID: %d, UserID: %d, Content: %s, Date: %s", s.ID, s.UserID, s.Content, s.Date)
}

type FavoriteReports struct {
	ID       int
	UserID   int
	ReportId int
}

type FinancialReport struct {
	ID              int
	Title           string
	ReportContentID int
	Date            time.Time
}

type ReportAnalysis struct {
	ID                int
	FinancialReportID int
	Result            string
}
