package covid

type SummaryResponse struct {
	Province map[string]int
	AgeGroup map[string]int
}

type CovidRepository interface{}

type CovidService interface{}
