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

	c := colly.NewCollector(
		colly.AllowedDomains("github.com", "www.github.com"),
	)

	c.OnHTML("div[id='user-repositories-list'] div h3 a", func(e *colly.HTMLElement) {
		repoLink := "https://github.com" + e.Attr("href")
		repositories = append(repositories, Repository{Link: repoLink})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping the page", r.Request.URL)

		file, _ := json.MarshalIndent(repositories, "", " ")

		_ = ioutil.WriteFile("repositories.json", file, 0644)
		fmt.Println("Saved results to repositories.json")
	})


	user := "ayushjain-10"
	err := c.Visit("https://github.com/" + user + "?tab=repositories")
	if err != nil {
		log.Fatal("Failed to visit: ", err)
	}
}
