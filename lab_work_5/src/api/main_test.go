package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"api/methods"
)

func TestGetAllFavoriteReports(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/favoriteReport", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestPostFavoriteReport(t *testing.T) {
	router := setupRouter()

	requestBody := map[string]int{
		"ReportID": 1,
	}
	bodyBytes, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "/favoriteReport", bytes.NewBuffer(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetFavoriteReportByID(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/favoriteReport/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestDeleteFavoriteReportByID(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("DELETE", "/favoriteReport/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestUpdateFavoriteReport(t *testing.T) {
	router := setupRouter()

	requestBody := map[string]int{
		"ID":       1,
		"ReportID": 2,
	}
	bodyBytes, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("PUT", "/favoriteReport", bytes.NewBuffer(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	methods.RegisterFavoriteReportRouts(router)

	return router
}
