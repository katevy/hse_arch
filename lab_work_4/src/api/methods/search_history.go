package methods

import (
	"api/database"
	"api/models"
	"api/sqlite"
	"api/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SearchHistoryResponse struct {
	ID      int
	Content string
	Date    time.Time
}

func RegisterSearchHistoryRouts(r *gin.Engine) {
	r.GET("/searchHistories", getAllSeacrhHistories)
	r.GET("/searchHistory/:id", getSearchHistoryByID)
	r.DELETE("/searchHistory/:id", deleteSearchHistoryByID)
	r.POST("/searchHistory", postSearchHistory)
	r.PUT("/searchHistory", updateSearchHistory)
}

func getAllSeacrhHistories(c *gin.Context) {

	db, err := sqlite.ConnectDB("database.db")

	if err != nil {
		log.Print("Failed to connect to database:", err)
	}

	defer db.DB.Close()

	result, err := database.SearchHistory.GetAllSeacrhHistory(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	var response []SearchHistoryResponse

	for _, history := range result {
		response = append(response, SearchHistoryResponse{
			ID:      history.ID,
			Content: history.Content,
			Date:    history.Date,
		})
	}

	c.IndentedJSON(http.StatusOK, response)
}

func postSearchHistory(c *gin.Context) {

	type SearchHistoryPostBody struct {
		Content string
		Date    string
	}

	var searchHistoryPostBody SearchHistoryPostBody

	var newSearchHistory models.SearchHistory

	if err := c.BindJSON(&searchHistoryPostBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if parsedDate, err := utils.ConvertBodyStringDateToTime(searchHistoryPostBody.Date); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		newSearchHistory.Date = parsedDate
		newSearchHistory.Content = searchHistoryPostBody.Content

		db, err := sqlite.ConnectDB("database.db")

		if err != nil {
			log.Print("Failed to connect to database:", err)
		}

		defer db.DB.Close()

		result, err := database.SearchHistory.CreateSearchHistory(db, &newSearchHistory)

		if err != nil {
			log.Print("Failed to create SearchHistory:", err)
			return
		}

		response := SearchHistoryResponse{
			ID:      result.ID,
			Content: result.Content,
			Date:    result.Date,
		}

		c.IndentedJSON(http.StatusOK, response)
	}

}

func getSearchHistoryByID(c *gin.Context) {
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

	result, err := database.SearchHistory.GetSearchHistoryByID(db, idInt)

	if err != nil {
		log.Printf("Failed to get SearchHistory by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get SearchHistory by ID"})
		return
	}

	response := SearchHistoryResponse{
		ID:      result.ID,
		Content: result.Content,
		Date:    result.Date,
	}

	c.IndentedJSON(http.StatusOK, response)
}

func deleteSearchHistoryByID(c *gin.Context) {
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

	deleteError := database.SearchHistory.DeleteSearchHistory(db, idInt)

	if deleteError != nil {
		log.Printf("Failed to delete SearchHistory by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete SearchHistory by ID"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"succes": "SearchHistory item deleted"})
}

func updateSearchHistory(c *gin.Context) {

	type SearchHistoryPutBody struct {
		ID      int
		Content string
		Date    string
	}

	var searchHistoryPutBody SearchHistoryPutBody

	var newSearchHistory models.SearchHistory

	if err := c.BindJSON(&searchHistoryPutBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if parsedDate, err := utils.ConvertBodyStringDateToTime(searchHistoryPutBody.Date); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		newSearchHistory.ID = searchHistoryPutBody.ID
		newSearchHistory.Date = parsedDate
		newSearchHistory.Content = searchHistoryPutBody.Content

		db, err := sqlite.ConnectDB("database.db")

		if err != nil {
			log.Print("Failed to connect to database:", err)
		}

		defer db.DB.Close()

		updateError := database.SearchHistory.UpdateSearchHistory(db, &newSearchHistory)

		if updateError != nil {
			log.Print("Failed to update SearchHistory:", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update SearchHistory by ID"})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"succes": "SearchHistory item updated"})
	}

}
