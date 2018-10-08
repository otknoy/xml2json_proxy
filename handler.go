package main

import (
	"fmt"
	"net/http"
)

// XML2JsonHandler is http handler that converts http response from XML to Json
type XML2JsonHandler struct {
}

func (h *XML2JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	url := "http://example.com"
	path := "/test"

	// build request

	// request

	// get response

	// convert response to json

	//

	res := "{\"foo\": 123}"

	fmt.Fprint(w, res)
}
