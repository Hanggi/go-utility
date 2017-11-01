package main

import (
	"flag"
	"fmt"

	"./spider"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("!!!")
	flag.Parse()

	// spider.HTTPGet()
	spider.ForExtract("https://cnodejs.org/?tab=all&page=", 2, func(i int, s *goquery.Selection) {
		title := s.Find(".topic_title").Text()

		fmt.Println(title)
	})
}
