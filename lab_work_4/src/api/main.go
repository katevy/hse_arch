package main

import (
	"github.com/gin-gonic/gin"

	"api/methods"
)

func main() {
	router := gin.Default()

	methods.RegisterSearchHistoryRouts(router)
	methods.RegisterFavoriteReportRouts(router)
	methods.RegisterFinancialReportRouts(router)

	router.Run("localhost: 8080")
}
