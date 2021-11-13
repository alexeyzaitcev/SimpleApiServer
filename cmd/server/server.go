package main

import (
	"fmt"
	"log"
	"net/http"

	myapi "github.com/alexeyzaitcev/SimpleApiServer/tree/master/internal/api/internal/api"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {

	myapi.Path = r.URL.Path[4:]
	myapi.Method = r.Method

	fmt.Println("API path: ", myapi.Path)
	fmt.Println("request method", myapi.Method)
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("form data: ", r.Form.Encode())

	// check URL
	switch apiPath := r.URL.Path[5:]; apiPath {
	case "methodA":
		fmt.Fprintf(w, "API method requested: %s!", apiPath)
	default:
		fmt.Fprintf(w, "API Method: '%s' not supported", apiPath)
	}

}

func main() {
	fmt.Println("Simple web server")

	fmt.Println(myapi.Version())

	// static content
	http.Handle("/", http.FileServer(http.Dir("./web/static")))

	// API server
	http.HandleFunc("/api/", apiHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
