package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserController(t *testing.T) {
	router := UserRoutes(gin.Default())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
