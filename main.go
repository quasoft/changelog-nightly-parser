// changelog-nightly-parser is an AWS Lambda function that visits the Changelog Nightly
// page, extracts the URLs of trending repositories found and stores them as a JSON file
// the the specified Github repository.
//
// Three environment variables have to be defined for the Lambda function to work:
// - GITHUB_REPOSITORY - name of repository to which to upload the JSON file (eg. "trending-daily").
// - GITHUB_OWNER - Github username (eg. "myusername")
// - GITHUB_TOKEN - Github personal token (eg. "myusername")
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/aws/aws-lambda-go/lambda"
)

// TrendingRepos is the structure used for marshaling the trending repositories to JSON.
// The three fields represent the three categories on Changelog's Nightly page:
// - First - repositories featured for the first time in the Changelog
// - New - new open sourced repositories
// - Repeaters - trending repos that have been featured before
type TrendingRepos struct {
	First     []string `json:"FirstTimers"`
	New       []string `json:"TopNew"`
	Repeaters []string `json:"RepeatPerformers"`
}

// The Downloader interface represent a type that can perform GET HTTP requests,
// like http.Client or StubDownloader.
type Downloader interface {
	Get(string) (*http.Response, error)
}

// The Uploader interface represent a type that can perform PUT HTTP requests,
// like http.Client or StubDownloader.
type Uploader interface {
	Do(*http.Request) (*http.Response, error)
}

var (
	downloader Downloader = &http.Client{}
	uploader   Uploader   = &http.Client{}
)

func download(t time.Time) (io.ReadCloser, error) {
	dateURL := "http://nightly.changelog.com/" + t.Format(`2006/01/02`)

	log.Printf("Getting data from %s", dateURL)
	resp, err := downloader.Get(dateURL)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// parseNightlyPage extracts all repository links from the ChangeLog's nightly page
// (http://nightly.changelog.com/YYYY/MM/DD/) in all three categories:
// - Top Starred Repositories – First Timers
// - Top New Repositories
// - Top Starred Repositories – Repeat Performers
func parseNightlyPage(body io.Reader) (*TrendingRepos, error) {
	trending := TrendingRepos{}

	doc, err := htmlquery.Parse(body)
	if err != nil {
		return nil, err
	}

	categoryClasses := []string{"top-all-firsts", "top-new", "top-all-repeats"}

	m := make(map[string][]string)
	for _, class := range categoryClasses {
		xpath := fmt.Sprintf(`//table[@id="%s"]//tr[contains(@class, 'about')]//a/@href`, class)
		for _, n := range htmlquery.Find(doc, xpath) {
			href := htmlquery.SelectAttr(n, "href")
			m[class] = append(m[class], href)
		}
	}

	trending.First = m[categoryClasses[0]]
	trending.New = m[categoryClasses[1]]
	trending.Repeaters = m[categoryClasses[2]]

	totalCount := len(trending.First) + len(trending.New) + len(trending.Repeaters)

	log.Printf("Found %d repositories", totalCount)

	return &trending, nil
}

func uploadToGithub(body []byte, path string, t time.Time) error {
	owner := os.Getenv("GITHUB_OWNER")
	if owner == "" {
		return fmt.Errorf("Upload GitHub owner not specified")
	}

	repo := os.Getenv("GITHUB_REPOSITORY")
	if repo == "" {
		return fmt.Errorf("Upload GitHub repository not specified")
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return fmt.Errorf("Upload GitHub token not specified")
	}

	// Create a file (https://developer.github.com/v3/repos/contents/#create-a-file):
	// This method creates a new file in a repository
	// PUT /repos/:owner/:repo/contents/:path
	u := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", owner, repo, path)

	params := struct {
		Message  string `json:"message"`
		Content  string `json:"content"`
		Commiter struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"committer"`
	}{}
	params.Message = "Uploading trending repos for " + t.Format("2006-01-02")
	params.Content = base64.StdEncoding.EncodeToString(body)
	params.Commiter.Name = "Bot"
	params.Commiter.Email = "bot@example.com"

	j, err := json.MarshalIndent(params, "", "    ")
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(j)

	r, err := http.NewRequest("PUT", u, buf)
	if err != nil {
		return err
	}
	r.Header.Set("Authorization", "token "+token)
	r.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := uploader.Do(r)
	if err != nil {
		return fmt.Errorf("uploading to %s failed with error: %v", u, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		msg, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			msg = []byte{}
		}
		return fmt.Errorf("uploading to %s failed with status %d, msg: %s", u, resp.StatusCode, string(msg))
	}

	return nil
}

// Handler is a lambda function that visits the Changelog Nightly page, extracts URLs
// to the trending repositories in all three categories, prepares a JSON file with the
// URLs and commits that file to a Github repository.
func Handler() error {
	// 1. Get HTML for current date
	yesterday := time.Now().AddDate(0, 0, -1)
	changelog, err := download(yesterday)
	if err != nil {
		return err
	}
	defer changelog.Close()

	// 2. Parse HTML and extract repository links
	trending, err := parseNightlyPage(changelog)
	if err != nil {
		return err
	}

	// 3. Build a JSON file with the links
	j, err := json.Marshal(trending)
	if err != nil {
		return err
	}

	// 4. Upload the file
	todaysFileName := yesterday.Format("2006-01-02.json")
	err = uploadToGithub(j, todaysFileName, yesterday)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
