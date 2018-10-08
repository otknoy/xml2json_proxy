package main

import (
	"net/http"
)

func main() {
	handler := &XML2JsonHandler{}

	http.HandleFunc("/", handler.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
