package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"
)

type CsvService struct{}

func NewCsvService() *CsvService {
	return &CsvService{}
}

func (cs *CsvService) ExtractData(file multipart.File) error {
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Ignorer les lignes vides (en-têtes et résumé) jusqu'à ce que nous trouvions une ligne avec des champs séparés par ';'
	// Si la ligne contient un seul champs, c'est un commentaire
	headers, err := cs.readHeaders(reader)
	if err != nil {
		return err
	}
	log.Println("Headers processed :", headers)

	// Lire les données du fichier
	data, err := cs.readBody(reader, headers)
	if err != nil {
		return err
	}
	log.Println("Body processed :", data)

	// Afficher les données
	for _, entry := range data {
		fmt.Println(entry)
	}
	return nil
}

func (*CsvService) readHeaders(reader *csv.Reader) ([]string, error) {
	var headers []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			// End of file
			break
		}
		if err != nil {
			fmt.Printf("\n\nHeader error : %v\n\n", err)
			return nil, err
		}

		if len(record) != 1 {
			headers = record
			break
		}
	}
	return headers, nil
}

func (*CsvService) readBody(reader *csv.Reader, headers []string) ([]map[string]string, error) {
	var data []map[string]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("\n\nBody error : %v\n\n", err)
			return nil, err
		}

		entry := make(map[string]string)
		for i, field := range record {
			key := headers[i]
			entry[key] = field
		}
		data = append(data, entry)
	}
	return data, nil
}
