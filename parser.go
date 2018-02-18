package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func parseRepository(parent *html.Node) (*Repository, error) {
	// Extract basic information about the repository
	a := htmlquery.FindOne(parent, `//tr[contains(@class, 'about')]//a`)
	if a == nil {
		// Ignore if URL or repository name cannot be determined
		return nil, fmt.Errorf("Could not determine repository URL and Name")
	}

	repo := Repository{
		URL:  htmlquery.SelectAttr(a, "href"),
		Name: htmlquery.InnerText(a),
	}

	p := htmlquery.FindOne(parent, `//tr[contains(@class, 'about')]//p`)
	if p != nil {
		repo.Description = strings.TrimSpace(htmlquery.InnerText(p))
	}

	s := htmlquery.FindOne(parent, `//span[contains(@title, 'Stars')]`)
	if s != nil {
		starsText := strings.TrimSpace(htmlquery.InnerText(s))
		sn, err := strconv.Atoi(starsText)
		if err == nil {
			repo.Stars = sn
		}
	}

	l := htmlquery.FindOne(parent, `//span[contains(@title, 'Language')]//a`)
	if l != nil {
		repo.Language = strings.TrimSpace(htmlquery.InnerText(l))
	}

	return &repo, nil
}

func parseCategory(parent *html.Node, category string) []Repository {
	list := []Repository{}

	filter := fmt.Sprintf(`//table[@id="%s"]//div[contains(@class, 'repository')]`, category)
	for _, n := range htmlquery.Find(parent, filter) {
		repo, err := parseRepository(n)
		if err != nil {
			continue
		}
		list = append(list, *repo)
	}

	return list
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

	trending.First = parseCategory(doc, "top-all-firsts")
	trending.New = parseCategory(doc, "top-new")
	trending.Repeaters = parseCategory(doc, "top-all-repeats")

	totalCount := len(trending.First) + len(trending.New) + len(trending.Repeaters)

	log.Printf("Found %d repositories", totalCount)

	return &trending, nil
}
