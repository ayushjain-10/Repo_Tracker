package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	var repos []Repository
	err = json.Unmarshal(body, &repos)
	if err != nil {
		fmt.Println(string(body))
		log.Fatal("Failed to unmarshal response body: ", err)
		return repositories
	}

	return repos
}



func scrapeGithubUser(user string) GithubUser {
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

	return githubUser
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

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("username")
		if user == "" {
			http.Error(w, "Username parameter is required", http.StatusBadRequest)
			return
		}
	
		githubUser := scrapeGithubUser(user)
	
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
