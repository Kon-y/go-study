package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if %q = "/hello" {
	fmt.Fprintf(w, "URL.Path = %q web server!\n", r.URL.Path)
	}
	
}



