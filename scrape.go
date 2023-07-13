package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"sync"
)

type Repository struct {
	Name   string `json:"name"`
	Link   string `json:"html_url"`
	Stars  int    `json:"stargazers_count"`
	Forks  int    `json:"forks_count"`
	Issues int    `json:"open_issues_count"`
}

type GithubUser struct {
	Name             string `json:"name"`
	RepoCount        int    `json:"public_repos"`
	StarredCount     int    `json:"starred_count"`
	FollowersCount   int    `json:"followers"`
	FollowingCount   int    `json:"following"`
}

var cache = NewCache()
var wg sync.WaitGroup

func scrapeGithub(user string, done chan<- bool) []Repository {

	if _, found := cache.Get(user); found {
		done <- true
	}

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

	var repos []Repository
	err = json.Unmarshal(body, &repos)
	if err != nil {
		fmt.Println(string(body))
		log.Fatal("Failed to unmarshal response body: ", err)
		return repositories
	}

	cache.Set(user, repos, 24*time.Hour)
	done <- true

	return repos
}



func scrapeGithubUser(user string, done chan <- bool) GithubUser {

	if _, found := cache.Get(user); found {
		done <- true
	}

	var githubUser GithubUser

	apiURL := fmt.Sprintf("https://api.github.com/users/%s", user)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		log.Fatal("Failed to create API request: ", err)
		return githubUser
	}

	req.Header.Set("Authorization", "Bearer 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")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Failed to retrieve user: ", err)
		return githubUser
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response body: ", err)
		return githubUser
	}

	err = json.Unmarshal(body, &githubUser)
	if err != nil {
		log.Fatal("Failed to unmarshal response body: ", err)
		return githubUser
	}

	cache.Set(user, githubUser, 24*time.Hour)
	done <- true

	return githubUser
}

func main() {
	http.HandleFunc("/repos", func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("username")
		if user == "" {
			http.Error(w, "Username parameter is required", http.StatusBadRequest)
			return
		}

		done := make(chan bool)
		go scrapeGithub(user, done) // start the goroutine

		<-done // wait for the goroutine to finish

		repositories, _ := cache.Get(user)

		js, err := json.Marshal(repositories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("username")
		if user == "" {
			http.Error(w, "Username parameter is required", http.StatusBadRequest)
			return
		}

		done := make(chan bool)
		go scrapeGithubUser(user, done) // start the goroutine

		<-done // wait for the goroutine to finish

		githubUser, _ := cache.Get(user)

		js, err := json.Marshal(githubUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.ListenAndServe(":8080", nil)
}