package adapters

import (
	"conso-tracker/src/external/views"
	"conso-tracker/src/internal/core/ports"
	"conso-tracker/src/internal/core/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type FileHandler struct {
	puissanceProcessor ports.PuissanceProcessor
}

func NewFileHandler() *FileHandler {
	return &FileHandler{services.NewPuissanceService()}
}

// RenderImportModal is the handler for the import modal template
func (fh *FileHandler) RenderImportModal(w http.ResponseWriter, r *http.Request) {
	modal := views.ImportModal()
	modal.Render(r.Context(), w)
}

func (fh *FileHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Read the uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	b, _ := json.Marshal(file)
	log.Printf("Uploaded file: %+v\n", string(b))

	// Process the file data
	chart, err := fh.puissanceProcessor.ProcessChart(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, _ = json.Marshal(*chart) // FIXME : handle error, process chart data
	fmt.Println("Chart is", string(b))

	// Send a response to indicate that the upload was successful
	w.Write([]byte("File uploaded successfully!"))
}
