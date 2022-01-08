package covid_test

import (
	"testing"

	"github.com/krastomer/covid-summary/internal/covid"
)

func TestGetCovidData(t *testing.T) {
	repository := covid.NewCovidRepository()
	_, err := repository.GetCovidData()
	if err != nil {
		t.Errorf("get error %s", err)
	}
}
