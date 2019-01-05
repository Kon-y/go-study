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

// stargaxersのURLを取得し、starまで取得してくる。
func GetStars(url string) {
	starsnum := 1
	doc, _ := goquery.NewDocument(url)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url_stars, _ := s.Attr("href")
		if starsnum <= 10 {
			if strings.Contains(url_stars, "stargazers") {

			}
			starsnum++
		}
	})
}

func main() {
	url := "https://github.com/trending"
	GetRank(url)
	GetStars(url)
}
