package main

import (
	"flag"

	"./p2"
)

func main() {

	// type Student struct {
	// 	Name    string
	// 	Age     int
	// 	Guake   bool
	// 	Classes []string
	// 	Price   float32
	// }

	// st := &Student{
	// 	"Xiao Ming",
	// 	16,
	// 	true,
	// 	[]string{"Math", "English", "Chinese"},
	// 	9.99,
	// }

	// b, err := json.Marshal(st)
	// fmt.Println(string(b))

	// fmt.Println("!!!")
	flag.Parse()

	p2.RockVmineSummaries()

	// spider.HTTPGet()
	// spider.ForExtract("https://cnodejs.org/?tab=all&page=", 2, func(i int, s *goquery.Selection) {
	// 	title := s.Find(".topic_title").Text()

	// 	fmt.Println(title)
	// })
}
