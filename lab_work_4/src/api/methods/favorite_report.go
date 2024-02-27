package methods

import (
	"api/database"
	"api/models"
	"api/sqlite"
	"api/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteReportResponse struct {
	ID       int
	ReportID int
}

func RegisterFavoriteReportRouts(r *gin.Engine) {
	r.GET("/favoriteReport", getAllfavoriteReports)
	r.GET("/favoriteReport/:id", getFavoriteReportByID)
	r.DELETE("/favoriteReport/:id", deleteFavoriteReportByID)
	r.POST("/favoriteReport", postFavoriteReport)
	r.PUT("/favoriteReport", updateFavoriteReport)
}

func getAllfavoriteReports(c *gin.Context) {

	db, err := sqlite.ConnectDB("database.db")

	if err != nil {
		log.Print("Failed to connect to database:", err)
	}

	defer db.DB.Close()

	result, err := database.FavoriteReport.GetAllFavoriteReports(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	var response []FavoriteReportResponse

	for _, fav_report := range result {
		response = append(response, FavoriteReportResponse{
			ID:       fav_report.ID,
			ReportID: fav_report.ReportID,
		})
	}

	c.IndentedJSON(http.StatusOK, response)
}

func postFavoriteReport(c *gin.Context) {

	type FavoriteReportPostBody struct {
		ReportID int
	}

	var favoriteReportPostBody FavoriteReportPostBody

	var newFavoriteReport models.FavoriteReport

	if err := c.BindJSON(&favoriteReportPostBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Print(favoriteReportPostBody)

	newFavoriteReport = models.FavoriteReport{
		ID:       0,
		UserID:   1,
		ReportID: favoriteReportPostBody.ReportID,
	}

	db, err := sqlite.ConnectDB("database.db")

	if err != nil {
		log.Print("Failed to connect to database:", err)
	}

	defer db.DB.Close()

	result, err := database.FavoriteReport.CreateFavoriteReport(db, &newFavoriteReport)

	if err != nil {
		log.Print("Failed to create FavoriteReport:", err)
		return
	}

	response := FavoriteReportResponse{
		ID:       result.ID,
		ReportID: result.ReportID,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func getFavoriteReportByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := utils.ConvertStringParamToIntegerID(id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := sqlite.ConnectDB("database.db")

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	defer db.DB.Close()

	result, err := database.FavoriteReport.GetFavoriteReportByID(db, idInt)

	if err != nil {
		log.Printf("Failed to get FavoriteReport by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get FavoriteReport by ID"})
		return
	}

	response := FavoriteReportResponse{
		ID:       result.ID,
		ReportID: result.ReportID,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func deleteFavoriteReportByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Failed to convert id to int:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	db, err := sqlite.ConnectDB("database.db")

	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	defer db.DB.Close()

	deleteError := database.FavoriteReport.DeleteFavoriteReport(db, idInt)

	if deleteError != nil {
		log.Printf("Failed to delete FavoriteReport by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete FavoriteReport by ID"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"succes": "FavoriteReport item deleted"})
}

func updateFavoriteReport(c *gin.Context) {

	type FavoriteReportPutBody struct {
		ID       int
		ReportID int
	}

	var favoriteReportPutBody FavoriteReportPutBody

	var newFavoriteReport models.FavoriteReport

	if err := c.BindJSON(&favoriteReportPutBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newFavoriteReport = models.FavoriteReport{
		ID:       favoriteReportPutBody.ID,
		UserID:   1,
		ReportID: favoriteReportPutBody.ReportID,
	}

	db, err := sqlite.ConnectDB("database.db")

	if err != nil {
		log.Print("Failed to connect to database:", err)
	}

	defer db.DB.Close()

	updateError := database.FavoriteReport.UpdateFavoriteReport(db, &newFavoriteReport)

	if updateError != nil {
		log.Print("Failed to update FavoriteReport:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update FavoriteReport by ID"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"succes": "FavoriteReport item updated"})
}
