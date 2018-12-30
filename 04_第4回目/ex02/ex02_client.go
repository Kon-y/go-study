package main

import (
	"log"
	"net/http"
)

func main() 
{
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.ListenAndServe(":9999", nil)

	res, error := http.Get("http://localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	Body, error := ioutil.ReadAll(res.Body)
	if error := nil{log.Fatal(error)}
    fmt.printf("%q\n", body[:])
}
