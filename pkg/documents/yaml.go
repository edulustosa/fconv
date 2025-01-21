package documents

import (
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

func ToYaml(file io.Reader, ext string) ([]byte, error) {
	switch ext {
	case "json":
		return jsonToYaml(file)
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
