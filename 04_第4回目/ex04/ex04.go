package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

//var lines int = 0

func set(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Println(url)
		//lines += 1
	})
}
/*
func main() {
	url := "https://github.com/trending"
	set(url)
	fmt.Printf("Number of [a] tags is %d\n", lines)
}
*/
