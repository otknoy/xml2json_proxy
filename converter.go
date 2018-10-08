package main

import (
	"io"

	xj "github.com/basgys/goxml2json"
)

// XML2JsonConverter converts XML to Json
type XML2JsonConverter struct {
}

func (c *XML2JsonConverter) convert(xml io.Reader) (string, error) {
	json, err := xj.Convert(xml)
	if err != nil {
		return "", err
	}

	return json.String(), nil
}
