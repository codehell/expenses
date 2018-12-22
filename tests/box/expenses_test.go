package box

import (
	"expenses/routes"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreateExpense(t *testing.T) {
	err := os.Setenv("GCP_ENVIRONMENT", "testing")
	if err != nil {
		log.Fatalln(err)
	}
	r := routes.Router()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/expenses", nil)
	r.ServeHTTP(w, req)
}
