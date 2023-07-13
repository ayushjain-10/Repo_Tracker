# Repo_Tracker [![Go Report Card](https://goreportcard.com/badge/github.com/ayushjain-10/repo_tracker-main)](https://goreportcard.com/report/github.com/ayushjain-10/repo_tracker-main) ![GitHub](https://img.shields.io/github/license/ayushjain-10/Repo_Tracker) 





##  🎥 Demo
<div align="center">
    <img src="output.gif" width="auto" height="540">
    <br>
</div>

## 📝 Description

The GitHub Repository Tracker is a program written in Go that tracks changes in GitHub repositories of the user by scraping the repository page. This tool is designed to help developers stay updated on repositories they are interested in or depend upon without having to constantly check GitHub manually.

It can notify team members about significant changes such as new releases, closed issues, and more.

## 🚀 Features

* **Track Multiple Repositories Concurrently:** By using Go's goroutines and channels, the tool can concurrently track multiple repositories, providing real-time updates.
* **Email Notifications:** Get instant email notifications whenever there's a significant update in any of the repositories being tracked.
* **Track :** Configure the tool to track stars, forks, releases, issues, and pull requests.
* Track **Multiple Users:** The tool can track multiple users and their repositories.

## Getting Started

To get a local copy up and running, follow these steps:

### Prerequisites

Ensure you have Go installed on your machine. If not, download and install [Go](https://golang.org/dl/).

### 📦 Installation

1. **Clone the repository**

   Open your terminal and run:

   ```sh
   git clone https://github.com/ayushjain-10/repo_tracker-main.git
    ```
   
2. **Navigate to the cloned repository**

   ```sh
   cd Repo_Tracker
    ```
   

3. **Setup environment variables** 
   Create a sendgrid.env file in the root directory and populate it with the following environment variables: SENDGRID_API_KEY and PRIVATE_API_KEY.

   Example:
   ```sh
   export SENDGRID_API_KEY='YOUR_API_KEY'
   export PRIVATE_API_KEY='YOUR_API_KEY'
   ```


3. **Install dependencies**

   ```sh
   go mod download
   ```

4. **Build the binary**

   ```sh
    go build
    ```

5. **Run the binary**

   ```sh
    go run .
    ```

### Frontend Setup

If you want to run the frontend of this application:

1. Open a new terminal and run:
    ```sh
    cd GitHubRepoTracker
    ```

2. Install the necessary npm packages:
    ```sh
    npm install
    ```

3. Start the application:
      ```sh
      npm run ios
      ```

## 👥 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

1. **Fork the Project**

2. **Create your Feature Branch**
   
   ```sh
   git checkout -b feature/YourFeature
   ```

3. **Commit your Changes**
   
   ```sh
   git commit -m 'Add some YourFeature'
   ```

4. **Push to the Branch**
   
   ```sh
   git push origin feature/YourFeature
   ```

5. **Open a Pull Request**

## 📝 License

Distributed under the MIT License. See `LICENSE` for more information.


## 🙌 Acknowledgements


- [Go Report Card](https://goreportcard.com/)

- [Github API](https://docs.github.com/en/rest/guides/getting-started-with-the-rest-api?apiVersion=2022-11-28)

- [SendGrid API](https://sendgrid.com/solutions/email-api/?utm_source=google&utm_medium=cpc&utm_term=sendgrid%20api&utm_campaign=SendGrid_G_S_NAMER_Brand_Tier1&cq_plac=&cq_net=g&cq_pos=&cq_med=&cq_plt=gp&gad=1&gclid=CjwKCAjwwb6lBhBJEiwAbuVUSnb6Ex5mWe5pdh0zVf7ngih0XIzI_L5PYgQBYFEbRXaCOLxX2lEDshoC8-wQAvD_BwE)

- [Golang](https://golang.org/)

- [React Native](https://reactnative.dev/)

- [Firebase](https://firebase.google.com/)


