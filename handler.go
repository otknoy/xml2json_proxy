package main

import (
	"fmt"
	"net/http"
)

// XML2JsonHandler is http handler that converts http response from XML to Json
type XML2JsonHandler struct {
}

func (h *XML2JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hoge")
}
