package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", http.StripPrefix("/", fs))

	go http.ListenAndServe(":9999", nil)

    func main() {
		moc, err := goquery.NewDocument("http://localhost:9999")
		if err != nil {
			fmt.Print("url scarapping failed")
		}
		doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			  url, _ := s.Attr("a")
			  fmt.Println(url)
		})
	}
}
