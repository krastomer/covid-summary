package covid

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type repository struct {
	client http.Client
}

var (
	ErrGetRequestError      = errors.New("get request error")
	ErrConvertToStructError = errors.New("convert to struct error")
)

func NewCovidRepository() CovidRepository {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	return &repository{client: client}
}

func (r *repository) GetCovidData() (data []Record, err error) {
	url := "https://static.wongnai.com/devinterview/covid-cases.json"

	response, err := r.client.Get(url)

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
