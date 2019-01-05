package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetRank(url string) {
	trend := 1
	doc, _ := goquery.NewDocument(url)
	doc.Find("li").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("id")
		if trend <= 10 {
			//　空行を出力させないため、len(url)を使用する。
			if len(url) > 0 {
				// リポ名の頭に「pa-」がつくので、削除する。
				replace := strings.Replace(url, "pa-", "", 1)
				fmt.Printf("Trend%d位は	→	%s\n", trend, replace)
				trend++
			}
		}
	})
}

func GetStars(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find("span").Each(func(_ int, s *goquery.Selection) {
		stars, _ := s.Attr("class")
		if len(stars) < 0 {
			fmt.Println(stars)
		}
	})
}

func main() {
	url := "https://github.com/trending"
	GetRank(url)
	GetStars(url)
	//fmt.Printf("URL: %s, Star: %s\n", url, urls)
	//	fmt.Printf("Number of [a] tags is %d\n", lines)
}
