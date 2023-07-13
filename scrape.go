package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Repository struct {
	Link   string `json:"link"`
	Stars  int    `json:"stars"`
	Forks  int    `json:"forks"`
	Issues int    `json:"issues"`
}

type GithubRepoResponse struct {
	StargazersCount int `json:"stargazers_count"`
	ForksCount      int `json:"forks_count"`
	OpenIssuesCount int `json:"open_issues_count"`
}

func fetchRepoDetails(repoURL string) (GithubRepoResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", repoURL, nil)
	if err != nil {
		return GithubRepoResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer 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")

	resp, err := client.Do(req)
	if err != nil {
		return GithubRepoResponse{}, err
	}
	defer resp.Body.Close()

	var repoResponse GithubRepoResponse
	err = json.NewDecoder(resp.Body).Decode(&repoResponse)
	if err != nil {
		return GithubRepoResponse{}, err
	}

	return repoResponse, nil
}

func scrapeGithub(user string) []Repository {
	var repositories []Repository

	apiURL := fmt.Sprintf("https://api.github.com/users/%s/repos", user)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatal("Failed to create API request: ", err)
		return repositories
	}

	req.Header.Set("Authorization", "Bearer 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")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Failed to retrieve user repositories: ", err)
		return repositories
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response body: ", err)
		return repositories
	}

	var repos []map[string]interface{}
	err = json.Unmarshal(body, &repos)
	if err != nil {
		log.Fatal("Failed to unmarshal response body: ", err)
		return repositories
	}

	for _, repo := range repos {
		repoLink := repo["html_url"].(string)
		repoDetails, err := fetchRepoDetails(repo["url"].(string))
		if err != nil {
			log.Printf("Failed to fetch details for repository %s: %v", repoLink, err)
			continue
		}

		repositories = append(repositories, Repository{
			Link:   repoLink,
			Stars:  repoDetails.StargazersCount,
			Forks:  repoDetails.ForksCount,
			Issues: repoDetails.OpenIssuesCount,
		})
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
