package xml_helper

import (
	"encoding/json"
	"fmt"
	"github.com/clbanning/mxj/v2"
)

func ConvertXMLToJSON(xmlString string) (string, error) {
	mv, err := mxj.NewMapXml([]byte(xmlString)) // Unmarshal XML to map
	if err != nil {
		return "", fmt.Errorf("error unmarshalling XML: %w", err)
	}

	jsonBytes, err := json.Marshal(mv) // Marshal map to JSON
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON: %w", err)
	}

	return string(jsonBytes), nil
}
