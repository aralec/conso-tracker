package services

import (
	linechart "conso-tracker/src/external/components/line-chart"
	"mime/multipart"
)

type PuissanceService struct {
	csvService *CsvService
}

func NewPuissanceService() *PuissanceService {
	return &PuissanceService{NewCsvService()}
}

func (ps *PuissanceService) ProcessChart(file multipart.File) (*linechart.Chart, error) {
	err := ps.csvService.ExtractData(file)
	if err != nil {
		return nil, err
	}

	return &linechart.Chart{}, nil
}

/*
	chart := components.Chart{
		ElementID: "chart",
		Data: components.Configuration{
			Type: "line",
			Data: components.Data{
				Labels: []string{"January", "February", "March", "April", "May", "June"},
				Datasets: []components.Dataset{
					{
						Label:       "My First dataset",
						Data:        []float64{65, 59, 80, 81, 56, 55, 40},
						Fill:        false,
						BorderColor: "rgb(75, 192, 192)",
						Tension:     0.1,
					},
				},
			},
		},
	}
*/
