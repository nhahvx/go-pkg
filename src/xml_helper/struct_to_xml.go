package xml_helper

import "encoding/xml"

func ConvertStructToXML(v interface{}) (string, error) {
	xmlData, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}

	return string(xmlData), nil
}
