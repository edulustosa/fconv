package documents

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"

	"github.com/sbabiv/xml2map"
	"gopkg.in/yaml.v3"
)

func ToYaml(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "json":
		return jsonToYaml(file)
	case "csv":
		return csvToYaml(file)
	case "xml":
		return xmlToYaml(file)
	}

	return nil, fmt.Errorf("unsupported file extension: %s to yaml", ext)
}

func jsonToYaml(file io.Reader) ([]byte, error) {
	var jsonDoc any
	if err := json.NewDecoder(file).Decode(&jsonDoc); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	yamlDoc, err := yaml.Marshal(jsonDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to yaml: %w", err)
	}

	return yamlDoc, nil
}

func csvToYaml(csvFile io.Reader) ([]byte, error) {
	csvReader := csv.NewReader(csvFile)
	headers, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv headers: %w", err)
	}

	var yamlContent []map[string]string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("failed to read csv record: %w", err)
		}

		yamlRecord := make(map[string]string)
		for i, header := range headers {
			yamlRecord[header] = record[i]
		}

		yamlContent = append(yamlContent, yamlRecord)
	}

	yamlBytes, err := yaml.Marshal(yamlContent)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to yaml: %w", err)
	}

	return yamlBytes, nil
}

func xmlToYaml(xmlFile io.Reader) ([]byte, error) {
	xmlDoc, err := xml2map.NewDecoder(xmlFile).Decode()
	if err != nil {
		return nil, fmt.Errorf("failed to read xml file: %w", err)
	}

	yamlDoc, err := yaml.Marshal(xmlDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to yaml: %w", err)
	}

	return yamlDoc, nil
}
