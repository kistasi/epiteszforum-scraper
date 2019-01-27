package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const (
	baseURL        = "http://epiteszforum.hu"
	numberOfPages  = 786
	articleLink    = "#archiveList article a"
	articleTitle   = "article#full hgroup h1"
	articleDate    = "article#full span.date"
	articleSummary = "article#full div#summary"
	articleContent = "article#full div#fullcontainer p"
)

func main() {
	interateOnPages()
}

func interateOnPages() {
	for page := 0; page < numberOfPages; page++ {
		iterateOnListingPage(buildListingURL(page))
	}
}

func buildListingURL(page int) string {
	return fmt.Sprintf("%s/archivum/0/0/0/%d", baseURL, page)
}

func buildArticleURL(articleURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, articleURL)
}

func getContentBySelector(collector *colly.HTMLElement, selector string) string {
	return strings.TrimSpace(collector.DOM.Find(selector).Text())
}

func iterateOnListingPage(url string) {
	listingCollector := colly.NewCollector()

	listingCollector.OnHTML(articleLink, func(article *colly.HTMLElement) {
		articleCollector := colly.NewCollector()

		articleCollector.OnHTML("body", func(articlePage *colly.HTMLElement) {
			url := buildArticleURL(article.Attr("href"))
			title := getContentBySelector(articlePage, articleTitle)
			date := getContentBySelector(articlePage, articleDate)
			summary := getContentBySelector(articlePage, articleSummary)
			content := getContentBySelector(articlePage, articleContent)

			fmt.Println("URL:", url)
			fmt.Println("Title:", title)
			fmt.Println("Date:", date)
			fmt.Println("Summary:", summary)
			fmt.Println("Content:", content)
			fmt.Println("===============")
		})

		articleCollector.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})

		err := articleCollector.Visit(buildArticleURL(article.Attr("href")))

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(2 * time.Second)
	})

	listingCollector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := listingCollector.Visit(url)

	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(2 * time.Second)
}
