package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func GetRank(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("li").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("id")
		fmt.Println(url)
	})
}

func GetStars(urls string) {
	doc, _ := goquery.NewDocument(urls)
	doc.Find("span").Each(func(_ int, s *goquery.Selection) {
		stars, _ := s.Attr("class")
		func main() {
			s := ""
			if len(s) == 0 {
			fmt.Println("空文字です")
			} else {
			fmt.Println("空文字ではありません")
			}
		}
		fmt.Println(stars)
	})
}

func main() {
	url := "https://github.com/trending"
	urls := ""
	GetRank(url)
	GetStars(urls)
	//fmt.Printf("URL: %s, Star: %s\n", url, urls)
	//	fmt.Printf("Number of [a] tags is %d\n", lines)
}
