package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	go http.ListenAndServe(":9999", nil)

	req, _ := http.NewRequest("GET", "http://localhost:9999", nil)
	//req, _ := http.NewRequest("GET", "http://www.rosso-tokyo.co.jp/", nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	html, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(html))
}
