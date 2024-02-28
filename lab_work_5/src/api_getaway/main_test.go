package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateWithAuthService(t *testing.T) {
	// Создаем фейковый сервер
	router := setupRouter()

	// Подготовка запроса с заголовком Authorization
	req, err := http.NewRequest("GET", "/proxy/service", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer validtoken")

	// Создаем ResponseRecorder (как фейковый клиент), чтобы записать ответ
	rr := httptest.NewRecorder()

	// Мы передаем наш запрос и ResponseRecorder на наш роутер, чтобы выполнить запрос
	router.ServeHTTP(rr, req)

	// Проверяем код статуса
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestProxyHandler(t *testing.T) {
	// Создаем фейковый сервер
	router := setupRouter()

	// Подготовка запроса
	req, err := http.NewRequest("GET", "/proxy/service", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder (как фейковый клиент), чтобы записать ответ
	rr := httptest.NewRecorder()

	// Мы передаем наш запрос и ResponseRecorder на наш роутер, чтобы выполнить запрос
	router.ServeHTTP(rr, req)

	// Проверяем код статуса
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

// Функция setupRouter создает экземпляр вашего маршрутизатора Gin и возвращает его
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Обработчики маршрутов
	router.GET("/proxy/:service", authenticateWithAuthService, proxyHandler)

	return router
}
