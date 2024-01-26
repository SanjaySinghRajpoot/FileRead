package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/SanjaySinghRajpoot/FileRead/controller"
	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetDataForFile(t *testing.T) {
	r := SetUpRouter()

	r.GET("/data", controller.GetData)

	reqFound, err := http.NewRequest(http.MethodGet, "http://localhost:8080/data?n=1", nil)

	if err != nil {
		fmt.Printf("test: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("working")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDataForLine(t *testing.T) {
	r := SetUpRouter()

	r.GET("/data", controller.GetData)

	reqFound, err := http.NewRequest(http.MethodGet, "http://localhost:8080/data?n=1&m=4", nil)

	if err != nil {
		fmt.Printf("test: error making http request: %s\n", err)
		os.Exit(1)
	}

	mockResponse := "test line"

	fmt.Println("working")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
