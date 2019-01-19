package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

const (
	baseURL = "http://epiteszforum.hu/"
	listingPage = baseURL + "archivum/0/0/0/1"
)

func main() {
		c := colly.NewCollector()

		c.OnHTML("#archiveList article a", func(e *colly.HTMLElement) {
			fmt.Println(e)
			c.Visit(fmt.Sprintf("%s%s", baseURL, e.Attr("href")))
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String());
		})

		c.Visit(listingPage)
}
