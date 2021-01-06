package main

import (
	"fmt"
	"net/http"
)

type APIHandler struct{}

func (APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.ListenAndServe(":8000", &APIHandler{})
}
