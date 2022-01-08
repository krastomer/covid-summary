package covid_test

import (
	"encoding/json"
	"testing"

	"github.com/krastomer/covid-summary/internal/covid"
)

type repository struct{}

func newMockRepository() covid.CovidRepository {
	return &repository{}
}

func (r *repository) GetCovidData() (data []covid.Record, err error) {
	jsonData := `
	{"Data":
		[
			{"ConfirmDate":"2021-05-04",
			"No":null,"Age":51,"Gender":"หญิง",
			"GenderEn":"Female","Nation":null,"NationEn":"China",
			"Province":"Phrae","ProvinceId":46,"District":null,"ProvinceEn":"Phrae","StatQuarantine":5},
			{"ConfirmDate":"2021-05-01","No":null,"Age":51,"Gender":"ชาย","GenderEn":"Male","Nation":null,
			"NationEn":"India","Province":"Suphan Buri","ProvinceId":65,"District":null,"ProvinceEn":"Suphan Buri",
			"StatQuarantine":8}
		]
	}
	`
	json.Unmarshal([]byte(jsonData), &data)

	return data, nil
}
func TestGetSummary(t *testing.T) {
	repo := newMockRepository()
	service := covid.NewCovidService(repo)

	_, err := service.GetSummary()
	if err != nil {
		t.Errorf("get error %s", err)
	}
}
