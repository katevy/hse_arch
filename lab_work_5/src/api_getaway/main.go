package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Маршрут для аутентификации и проксирования запросов
	router.Any("/proxy/:service", authenticateWithAuthService, proxyHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func authenticateWithAuthService(c *gin.Context) {
	// Извлечение токена из заголовка запроса
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}

	// parts := strings.Split(tokenString, " ")
	// log.Print(parts)
	// if len(parts) != 2 || parts[0] != "Bearer" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
	// 	c.Abort()
	// 	return
	// }

	// Отправка токена на ваш сервис авторизации для проверки
	url := "http://auth:8082/auth"
	req, err := http.NewRequest("GET", url, nil)
	log.Print(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate token"})
		c.Abort()
		return
	}

	// Установка заголовка Authorization с токеном
	req.Header.Set("Authorization", tokenString)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate token"})
		c.Abort()
		return
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}

func proxyHandler(c *gin.Context) {
	// Здесь вы можете реализовать логику проксирования запроса к другим сервисам
	service := c.Param("service")
	log.Print(service)
	resp, err := http.Get("http://server:8083/" + service)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to proxy request"})
		return
	}
	defer resp.Body.Close()

	// Копируем заголовки ответа
	for k, vv := range resp.Header {
		for _, v := range vv {
			c.Writer.Header().Add(k, v)
		}
	}

	// Копируем статус ответа
	c.Writer.WriteHeader(resp.StatusCode)

	// Копируем тело ответа
	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to proxy response"})
		return
	}
}
