package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//url := "http://golang.jp/"
	inputurl := flag.String("url", "http://localhost:8000", "Type your URL without http://")

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray)) // htmlをstringで取得
}
