package tests

import (
	"expenses/routes"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusUnautorizedWithoutToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	e := gin.Default()
	r := routes.Router(e)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/expenses", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
