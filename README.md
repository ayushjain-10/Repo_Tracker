# Repo_Tracker [![Go Report Card](https://goreportcard.com/badge/github.com/ayushjain-10/repo_tracker-main)](https://goreportcard.com/report/github.com/ayushjain-10/repo_tracker-main)

## Description

The GitHub Repository Tracker is a program written in Go that tracks changes in GitHub repositories of the user by scraping the repository page. This tool is designed to help developers stay updated on repositories they are interested in or depend upon without having to constantly check GitHub manually.

It can notify team members about significant changes such as new releases, closed issues, and more.

## Features

1. **Track Multiple Repositories Concurrently:** By using Go's goroutines and channels, the tool can concurrently track multiple repositories, providing real-time updates.
2. **Email Notifications:** Get instant email notifications whenever there's a significant update in any of the repositories being tracked.
3. **Filter Notifications:** Configure the tool to send notifications based on specific events like new releases, closed issues, etc.
4. **Flexible and Configurable:** You can add or remove repositories from the tracking list at any time.
5. **Slack Integration:** Integrate with Slack to receive notifications on a specified Slack channel.

## Getting Started

### Installation

1. Clone the repository

   ```sh
   git clone
    ```
2. Install dependencies

   ```sh
   go mod download
   ```

3. Build the binary

   ```sh
    go build
    ```
4. Run the binary

   ```sh
    go run *.go
    ```
5. To Run Frontend:
    Open a new terminal and run:
    ```sh
    npm install
    ```
    and
      ```sh
      npm run ios
      ```




