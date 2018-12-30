package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(":9999", nil)

	/* GetメソッドでURLにアクセス */
	res, err := http.get("http://localhost:9999")
	if err != nil {
		log.Fatal(err)
	}
}
