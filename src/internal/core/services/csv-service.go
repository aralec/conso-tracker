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

	// Ignorer les lignes vides (en-têtes et résumé) jusqu'à ce que nous trouvions une ligne avec des champs séparés par ';'
	// Si la ligne contient un seul champs, c'est un commentaire
	headers, abort := cs.readHeaders(reader)
	if abort {
		return
	}

	// Lire les données du fichier
	data, abort := cs.readBody(reader, headers)
	if abort {
		return
	}

	// Afficher les données
	for _, entry := range data {
		fmt.Println(entry)
	}
}

func (*CsvService) readHeaders(reader *csv.Reader) ([]string, bool) {
	var headers []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Header error :", err)
			return nil, true
		}

		if len(record) != 1 {
			headers = record
			break
		}
	}
	return headers, false
}

func (*CsvService) readBody(reader *csv.Reader, headers []string) ([]map[string]string, bool) {
	var data []map[string]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Body error :", err)
			return nil, true
		}

		entry := make(map[string]string)
		for i, field := range record {
			key := headers[i]
			entry[key] = field
		}
		data = append(data, entry)
	}
	return data, false
}
