package main

import (
	"fmt"
	"log"
	"net/http"
)

//homepageData prints a message to given I/O
func homepageData(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "Homepage endpoint has been hit with request.")
}

func requestHandler() {

	// Map requests for the root to the homepageData function
	http.HandleFunc("/", homepageData)

	// Log any errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	requestHandler()
}
