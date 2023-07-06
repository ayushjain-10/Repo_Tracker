package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: github.com, www.github.com
		colly.AllowedDomains("github.com", "www.github.com"),
	)

	// Callback for when visiting a repository page
	c.OnHTML("div[id='user-repositories-list'] div h3 a", func(e *colly.HTMLElement) {
		repoLink := e.Attr("href")
		fmt.Printf("Found repository: https://github.com%s\n", repoLink)
	})

	// Callback for handling scraped data
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping the page", r.Request.URL)
	})

	// Start the scraping process
	user := "ayushjain-10" // replace with the username you want to scrape
	err := c.Visit("https://github.com/" + user + "?tab=repositories")
	if err != nil {
		log.Fatal("Failed to visit: ", err)
	}
}
