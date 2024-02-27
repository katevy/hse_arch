package main

import (
	"github.com/gin-gonic/gin"

	"api/methods"
)

func main() {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	methods.RegisterSearchHistoryRouts(router)
	methods.RegisterFavoriteReportRouts(router)
	methods.RegisterFinancialReportRouts(router)

	router.Run(":8083")
}
