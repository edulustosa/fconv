package documents

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

func ToJson(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "csv":
		return csvToJson(file)
	case "yaml", "yml":
		return yamlToJson(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s to json", ext)
}

func csvToJson(file io.Reader) ([]byte, error) {
	csvReader := csv.NewReader(file)
	headers, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv headers: %w", err)
	}

	var jsonContent []map[string]string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("failed to read csv record: %w", err)
		}

		jsonRecord := make(map[string]string)
		for i, header := range headers {
			jsonRecord[header] = record[i]
		}

		jsonContent = append(jsonContent, jsonRecord)
	}

	jsonBytes, err := json.MarshalIndent(jsonContent, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to convert to json: %w", err)
	}

	return jsonBytes, nil
}

func yamlToJson(file io.Reader) ([]byte, error) {
	var yamlDoc map[string]any
	if err := yaml.NewDecoder(file).Decode(&yamlDoc); err != nil {
		return nil, fmt.Errorf("failed to decode yaml: %w", err)
	}

	jsonBytes, err := json.MarshalIndent(yamlDoc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to convert to json: %w", err)
	}

	return jsonBytes, nil
}
