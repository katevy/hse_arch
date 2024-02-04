package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"api/database"
	"api/models"
	"api/sqlite"
	"api/utils"
	"log"
)

func main() {
	var err error

	db, err := sqlite.ConnectDB("database.db")

	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	fmt.Println(db)

	router := gin.Default()
	router.GET("/searchHistories", getAllSeacrhHistories)
	router.GET("/searchHistory/:id", getSearchHistoryByID)
	router.DELETE("/searchHistory/:id", deleteSearchHistoryByID)
	router.POST("/searchHistory", postSearchHistory)

	router.Run("localhost: 8080")
}

func getAllSeacrhHistories(c *gin.Context) {

	db, err := sqlite.ConnectDB("database.db")

	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	defer db.DB.Close()

	result, err := database.SearchHistory.GetAllSeacrhHistory(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
	}

	c.IndentedJSON(http.StatusOK, result)
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
			log.Fatalln("Failed to connect to database:", err)
		}

		defer db.DB.Close()

		result, err := database.SearchHistory.CreateSearchHistory(db, &newSearchHistory)

		if err != nil {
			log.Fatalln("Failed to create SearchHistory:", err)
			return
		}

		c.IndentedJSON(http.StatusOK, result)
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

	fmt.Println(idInt)
	if err != nil {
		log.Printf("Failed to get SearchHistory by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get SearchHistory by ID"})
		return
	}

	c.IndentedJSON(http.StatusOK, result)
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
