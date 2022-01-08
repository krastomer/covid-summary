package infrastructure_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/krastomer/covid-summary/internal/infrastructure"
)

func TestCovidRouter(t *testing.T) {
	router := infrastructure.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/covid/summary", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("get error %s", w.Body.String())
	}
}
