package rendering

import (
	components "conso-tracker/src/external/components/line-chart"
	"conso-tracker/src/external/views"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
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

	home := views.Home(chart)
	home.Render(r.Context(), w)
}
