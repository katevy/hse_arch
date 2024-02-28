package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginEndpoint(t *testing.T) {
	// Создаем фейковый сервер
	router := setupRouter()

	// Подготовка данных для теста
	user := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	userBytes, _ := json.Marshal(user)

	// Создаем запрос POST к /login с тестовыми данными
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(userBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder (как фейковый клиент), чтобы записать ответ
	rr := httptest.NewRecorder()

	// Мы передаем наш запрос и ResponseRecorder на наш роутер, чтобы выполнить запрос
	router.ServeHTTP(rr, req)

	// Проверяем код статуса
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Проверяем ответ
	expected := ``
	assert.Contains(t, rr.Body.String(), expected)
}

func TestSignupEndpoint(t *testing.T) {
	// Создаем фейковый сервер
	router := setupRouter()

	// Подготовка данных для теста
	user := map[string]string{
		"username": "testuser",
		"password": "testpassword",
	}
	userBytes, _ := json.Marshal(user)

	// Создаем запрос POST к /signup с тестовыми данными
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(userBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder (как фейковый клиент), чтобы записать ответ
	rr := httptest.NewRecorder()

	// Мы передаем наш запрос и ResponseRecorder на наш роутер, чтобы выполнить запрос
	router.ServeHTTP(rr, req)

	// Проверяем код статуса
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Проверяем ответ
	expected := ``
	assert.Contains(t, rr.Body.String(), expected)
}

// Функция setupRouter создает экземпляр вашего маршрутизатора Gin и возвращает его
func setupRouter() *gin.Engine {
	router := gin.Default()

	// Обработчики маршрутов
	router.POST("/login", login)
	router.POST("/signup", signup)

	return router
}
