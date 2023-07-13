package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

type Repository struct {
	Link string `json:"link"`
}

func scrapeGithub(user string) []Repository {
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
	})

	err := c.Visit("https://github.com/" + user + "?tab=repositories")
	if err != nil {
		log.Fatal("Failed to visit: ", err)
	}

	return repositories
}

func main() {
	http.HandleFunc("/repos", func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("username")
		if user == "" {
			http.Error(w, "Username parameter is required", http.StatusBadRequest)
			return
		}

		repositories := scrapeGithub(user)

		js, err := json.Marshal(repositories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.ListenAndServe(":8080", nil)
}
