package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
)

type CsvService struct{}

func NewCsvService() *CsvService {
	return &CsvService{}
}

func (cs *CsvService) ExtractData(file multipart.File) {
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// En-têtes du fichier :
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Error while reading file :", err)
		return
	}

	// Afficher les en-têtes
	fmt.Println("\n\nHeaders are :\n", headers)

	// Lire les données du fichier
	var data []map[string]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			// Fin de ligne
			break
		}
		if err != nil {
			fmt.Println("Error while reading line : ", err)
			return
		}

		entry := make(map[string]string)
		for i, field := range record {
			key := headers[i]
			entry[key] = field
		}
		data = append(data, entry)
	}

	// Afficher les données
	for _, entry := range data {
		fmt.Println(entry)
	}
}
