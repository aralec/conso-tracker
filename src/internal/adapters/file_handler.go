package adapters

import (
	"conso-tracker/src/external/views"
	"fmt"
	"net/http"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

// RenderImportModal is the handler for the import modal template
func (fh *FileHandler) RenderImportModal(w http.ResponseWriter, r *http.Request) {
	modal := views.ImportModal()
	modal.Render(r.Context(), w)
}

func (fh *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Read the uploaded file
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Process the file data (e.g., store it in a database or perform some calculation)
	fmt.Printf("\n\nHeaders: %v\n\n", header)
	fmt.Printf("\n\nFile uploaded: %v\n\n", file)

	// Send a response to indicate that the upload was successful
	w.Write([]byte("File uploaded successfully!"))
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
