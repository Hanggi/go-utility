package spider

import (
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var (
// index int = 1
)

// Option vv
type Option struct {
	url      string
	urlParam string
	cpus     int
}

// HTTPGet vv
// func HTTPGet() {
// 	// resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
// 	resp, err := http.Get("http://blog.hanggi.me")
// 	if err != nil {
// 		// handle error
// 	}

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		// handle error
// 	}

// 	fmt.Println(string(body))

// 	doc, err := goquery.NewDocument("https://blog.hanggi.me")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Find the review items
// 	doc.Find("body").Each(func(i int, s *goquery.Selection) {
// 		// For each item found, get the band and title
// 		band := s.Find("p").Text()
// 		title := s.Find("h2").Text()
// 		fmt.Printf("Review %d: %s - %s\n", i, band, title)
// 	})
// }

// ForExtract vv
func ForExtract(url string, times int, query func(i int, s *goquery.Selection)) {
	index := 0
	if times < 1 {
		index = 1
	}
	for ; index <= times; index++ {
		doc, err := goquery.NewDocument(url + strconv.Itoa(index))
		if err != nil {
			log.Fatal(err)
		}
		doc.Find("body").Each(query)

		// updateURL()
	}
}
