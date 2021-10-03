package main

import (
	"fmt"
	"log"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {

	// check URL
	switch apiMethod := r.URL.Path[8:]; apiMethod {
	case "methodA":
		fmt.Fprintf(w, "API method requested: %s!", apiMethod)
	default:
		fmt.Fprintf(w, "API Method: '%s' not supported", apiMethod)
	}

}

func main() {
	fmt.Println("Simple web server")

	// static content
	http.Handle("/", http.FileServer(http.Dir("./web/static")))

	// API server
	http.HandleFunc("/api/", apiHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
