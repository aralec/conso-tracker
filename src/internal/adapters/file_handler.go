package adapters

import (
	"conso-tracker/src/external/components/modal"
	"net/http"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

// Importer TODO : importer un .csv
func (fh *FileHandler) Importer(w http.ResponseWriter, r *http.Request) {
	modal := modal.NewModal("Test")
	modal.Component().Render(r.Context(), w)
}

// TODO : Créer un service de transformation des données en chart
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
