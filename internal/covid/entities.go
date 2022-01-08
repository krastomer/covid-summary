package covid

type SummaryResponse struct {
	Province map[string]int
	AgeGroup map[string]int
}

type ResponseRecords struct {
	Data []Record
}

type Record struct {
	ConfirmDate    *string
	No             *int
	Age            *int
	Gender         *string
	GenderEn       *string
	Nation         *string
	NationEn       *string
	Province       *string
	ProvinceId     *int
	District       *string
	ProvinceEn     *string
	StatQuarantine *int
}

type CovidRepository interface {
	GetCovidData() (data []Record, err error)
}

type CovidService interface {
	GetSummary() (response SummaryResponse, err error)
}
