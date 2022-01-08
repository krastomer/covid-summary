package covid

type service struct {
	repo CovidRepository
}

func NewCovidService(repo CovidRepository) CovidService {
	return &service{repo: repo}
}

func (s *service) GetSummary() (response SummaryResponse, err error) {
	data, err := s.repo.GetCovidData()
	if err != nil {
		return response, err
	}

	provinceCount := make(map[string]int)
	ageCount := make(map[string]int)
	for _, row := range data {
		if row.ProvinceEn != nil {
			provinceCount[*row.ProvinceEn] += 1
		}
		switch s.findRangeAge(row.Age) {
		case 0:
			ageCount["N/A"] += 1
		case 1:
			ageCount["0-30"] += 1
		case 2:
			ageCount["31-60"] += 1
		case 3:
			ageCount["61+"] += 1
		}
	}

	response.Province = provinceCount
	response.AgeGroup = ageCount

	return response, nil
}

func (s *service) findRangeAge(age *int) (rowIn int) {
	if age == nil {
		return 0
	}
	if *age >= 0 && *age <= 30 {
		return 1
	}
	if *age >= 31 && *age <= 60 {
		return 2
	}
	return 3
}
