package covid

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type repository struct {
}

var (
	ErrGetRequestError      = errors.New("get request error")
	ErrConvertToStructError = errors.New("convert to struct error")
)

func NewCovidRepository() CovidRepository {
	return &repository{}
}

func (r *repository) GetCovidData() (data []Record, err error) {
	url := "http://static.wongnai.com/devinterview/covid-cases.json"

	response, err := http.Get(url)
	if err != nil {
		return nil, ErrGetRequestError
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	var records ResponseRecords
	if err = json.Unmarshal(body, &records); err != nil {
		return nil, ErrConvertToStructError
	}

	return records.Data, nil
}
