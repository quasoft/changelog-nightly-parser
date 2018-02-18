package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// Repository contains fields for the most relevant information available for each repository.
type Repository struct {
	Name        string `json:"Name"`
	URL         string `json:"URL"`
	Description string `json:"Description"`
	Stars       int    `json:"Stars"`
	Language    string `json:"Language"`
	Screenshot  string `json:"Screenshot"`
}

// TrendingRepos is the structure used for marshaling the trending repositories to JSON.
// The three fields represent the three categories on Changelog's Nightly page:
// - First - repositories featured for the first time in the Changelog
// - New - new open sourced repositories
// - Repeaters - trending repos that have been featured before
type TrendingRepos struct {
	First     []Repository `json:"FirstTimers"`
	New       []Repository `json:"TopNew"`
	Repeaters []Repository `json:"RepeatPerformers"`
}

// readmeURL() return the URL to the default readme of the repository
// (eg. https://api.github.com/repos/user1/repo1/readme).
func (r *Repository) readmeURL() string {
	u := strings.Replace(r.URL, "www.github.com", "api.github.com/repos", 1)
	u = strings.Replace(u, "/github.com", "/api.github.com/repos", 1)
	u = strings.TrimRight(u, "/") + "/readme"
	return u
}

// rawImageURL() returns the absolute URL to an image hosted inside a repository,
// given the branch name and the relative path to the image.
// (eg. https://raw.githubusercontent.com/user1/repo1/master/screenshot.jpg).
func (r *Repository) rawImageURL(branch string, relativePath string) string {
	u := strings.Replace(r.URL, "www.github.com", "raw.githubusercontent.com", 1)
	u = strings.Replace(u, "/github.com", "/raw.githubusercontent.com", 1)
	u = strings.TrimRight(u, "/") + "/" + branch + "/" + relativePath
	return u
}

// getReadmeHTML() performs a GET request to the Github API, retrieving the HTML
// of the default/main readme file in the repository.
// Parses the HTML and returns a pointer to the root html.Node of the document.
func (r *Repository) getReadmeHTML() (*html.Node, error) {
	req, err := http.NewRequest("GET", r.readmeURL(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.html")
	resp, err := downloader.Do(req)
	if err != nil {
		log.Printf("GET request for readme HTML failed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	root, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return root, nil
}

// populateScreenshot() downloads the default readme of the repository, finds the
// first image that appears to be a screenshot and populates the Screenshot field
// with the absolute URL to that image.
func (r *Repository) populateScreenshot() error {
	// Download the default readme file
	root, err := r.getReadmeHTML()
	if err != nil {
		log.Printf("Could not get repository readme file, error: %v", err)
		return err
	}

	// Find a screenshot in the readme file
	absURL := screenshotFromHTML(root)
	if absURL == "" {
		log.Printf("No screenshot detected for %s", r.URL)
		return fmt.Errorf("No screenshot detected")
	}

	if !strings.HasPrefix(absURL, strings.ToLower("http")) {
		// If a relative URL was found, use the githib repository as a base URL
		// TODO: Don't assume the branch is "master", get the branch name from Github API
		absURL = r.rawImageURL("master", absURL)
	}
	r.Screenshot = absURL
	log.Printf("Screenshot chosen for %s: %s", r.URL, r.Screenshot)

	return nil
}

// populateScreenshots() executes populateScreenshot() on the repositories in all three
// categories inside TrendingRepos.
func (tr *TrendingRepos) populateScreenshots() {
	log.Print("Populating screenshots")

	for i, r := range tr.First {
		r.populateScreenshot()
		tr.First[i] = r
	}

	for i, r := range tr.New {
		r.populateScreenshot()
		tr.New[i] = r
	}

	for i, r := range tr.Repeaters {
		r.populateScreenshot()
		tr.Repeaters[i] = r
	}
}
