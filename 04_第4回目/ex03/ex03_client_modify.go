package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

//　URLを指定せず、実行する場合
func cli() {
	urlarg := flag.String("url", "https://github.com/n-guitar/go_study", "まじか・・・")
	flag.Parse()
	req, _ := http.NewRequest("GET", "http://localhost:9999", nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	html, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(html))
}

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	go http.ListenAndServe(":9999", nil)
	go cli()
}
