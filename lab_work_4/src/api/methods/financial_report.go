package methods

import (
	"api/database"
	"api/models"
	"api/sqlite"
	"api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterFinancialReportRouts(r *gin.Engine) {

	r.POST("/financialReport", postFinancialReport)

}

func postFinancialReport(c *gin.Context) {

	type FinancialReportPostBody struct {
		Title   string
		Content string
		Date    string
	}

	var financialReportPostBody FinancialReportPostBody

	var newFinancialReport models.FinancialReport

	if err := c.BindJSON(&financialReportPostBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if parsedDate, err := utils.ConvertBodyStringDateToTime(financialReportPostBody.Date); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {

		newFinancialReport = models.FinancialReport{
			Title:   financialReportPostBody.Title,
			Content: financialReportPostBody.Content,
			Date:    parsedDate,
		}

		db, err := sqlite.ConnectDB("database.db")

		if err != nil {
			log.Print("Failed to connect to database:", err)
		}

		defer db.DB.Close()

		result, err := database.FinancialReport.CreateFinancialReport(db, &newFinancialReport)

		if err != nil {
			log.Print("Failed to create FinancialReport:", err)
			return
		}

		c.IndentedJSON(http.StatusOK, result)
	}

}
