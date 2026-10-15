package main

import (
	"fmt"
	"net/http"
)

var PORT string = ":8080"

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(PORT, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my API")
}