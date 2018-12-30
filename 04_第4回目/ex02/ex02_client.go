package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	go http.ListenAndServe(":9999", nil)

	res, error := http.Get("http://localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	Body, err := ioutil.ReadAll(res.Body)
	if err := nil {
        log.Fatal(err)
    }
	
	fmt.printf("%q\n", Body[:])
}
