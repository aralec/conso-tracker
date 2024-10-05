package adapters

import (
	"conso-tracker/src/external/views"
	"conso-tracker/src/internal/core/ports"
	"conso-tracker/src/internal/core/services"
	"fmt"
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

	// Process the file data
	chart := fh.puissanceProcessor.ProcessChart(file)
	fmt.Println("Chart is", chart)

	// Send a response to indicate that the upload was successful
	w.Write([]byte("File uploaded successfully!"))
}
