package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request came through to Home Page")
}
