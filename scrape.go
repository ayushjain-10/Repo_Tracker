package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

type Repository struct {
	Link string `json:"link"`
}

func main() {
	var repositories []Repository

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: github.com, www.github.com
		colly.AllowedDomains("github.com", "www.github.com"),
	)

	// Callback for when visiting a repository page
	c.OnHTML("div[id='user-repositories-list'] div h3 a", func(e *colly.HTMLElement) {
		repoLink := "https://github.com" + e.Attr("href")
		repositories = append(repositories, Repository{Link: repoLink})
	})

	// Callback for handling scraped data
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping the page", r.Request.URL)

		// After finished scraping, save our results to a .json file
		file, _ := json.MarshalIndent(repositories, "", " ")

		_ = ioutil.WriteFile("repositories.json", file, 0644)
		fmt.Println("Saved results to repositories.json")
	})

	// Start the scraping process
	user := "ayushjain-10" // replace with the username you want to scrape
	err := c.Visit("https://github.com/" + user + "?tab=repositories")
	if err != nil {
		log.Fatal("Failed to visit: ", err)
	}
}
