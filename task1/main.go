package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	ForkNumber  int    `json:"forks_count"`
	CreatedAt   string `json:"created_at"`
	Link        string `json:"html_url"`
}

func getRepo(owner, repoName string) (*Repo, error) {
	gh, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repoName))
	if err != nil {
		return nil, fmt.Errorf("No network connection")
	}

	defer gh.Body.Close()

	if gh.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("repository %s/%s not found", owner, repoName)
	}

	if gh.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status code: %d", gh.StatusCode)
	}

	body, err := io.ReadAll(gh.Body)
	if err != nil {
		return nil, err
	}

	var repo Repo
	err = json.Unmarshal(body, &repo)
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (r *Repo) String() string {
	createdAt := r.CreatedAt
	if t, err := time.Parse(time.RFC3339, r.CreatedAt); err == nil {
		createdAt = t.Format("January 2, 2006")
	}
	return fmt.Sprintf("Name: %s\nDescription: %s\nStars: %d\nForks: %d\nCreated at: %s\nLink: %s\n",
		r.Name, r.Description, r.Stars, r.ForkNumber, createdAt, r.Link)
}

func parseInput(input string) (owner, repo string, ok bool) {
	input = strings.TrimSpace(input)
	input = strings.TrimPrefix(input, "https://")
	input = strings.TrimPrefix(input, "http://")
	input = strings.TrimPrefix(input, "github.com/")
	input = strings.TrimSuffix(input, "/")

	parts := strings.SplitN(input, "/", 2)
	if len(parts) == 2 && parts[0] != "" && parts[1] != "" {
		return parts[0], parts[1], true
	}
	return "", "", false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <github-url-or-owner/repo>")
		fmt.Println("Example: go run main.go golang/go")
		os.Exit(1)
	}

	input := os.Args[1]

	owner, repoName, ok := parseInput(input)
	if !ok {
		fmt.Println("Invalid input. Use 'owner/repo' or a GitHub URL.")
		return
	}

	repo, err := getRepo(owner, repoName)
	if err != nil {
		fmt.Printf("Error getting repo: %v\n", err)
		return
	}
	fmt.Println(repo)
}
