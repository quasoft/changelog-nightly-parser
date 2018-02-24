package main

import (
	"bytes"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func screenshotFromHTML(parent *html.Node) string {
	images := htmlquery.Find(parent, `//img`)
	screenshot := screenshotFromImages(images)
	return screenshot
}

// screenshotFromImages() returns the first image that looks like a screenshot
// and returns the "src" attribute of that image as-is.
func screenshotFromImages(images []*html.Node) string {
	var src string
	var screenshot *html.Node

	for _, img := range images {
		outerHTML, err := nodeToHTML(img)
		if err != nil {
			continue
		}

		outerHTML = strings.ToLower(outerHTML)
		if isBadge(outerHTML) || isIcon(outerHTML) || isLogo(outerHTML) {
			continue
		}

		// Remember the first image that is not a badge or icon.
		// If no screenshot is found, return this.
		if screenshot == nil {
			screenshot = img
		}

		if strings.Contains(outerHTML, "screen") ||
			strings.Contains(outerHTML, "demo") ||
			strings.Contains(outerHTML, "example") ||
			strings.Contains(outerHTML, "sample") {
			screenshot = img
			break
		}
	}

	if screenshot != nil {
		src = htmlquery.SelectAttr(screenshot, "src")
	}
	return src
}

func nodeToHTML(node *html.Node) (string, error) {
	var buf bytes.Buffer
	err := html.Render(&buf, node)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func isBadge(html string) bool {
	return strings.Contains(html, "badge") ||
		strings.Contains(html, "shields.io") ||
		strings.Contains(html, "travis-ci.org") ||
		strings.Contains(html, "coveralls.io") ||
		strings.Contains(html, "snyk.io") ||
		strings.Contains(html, "david-dm.org") ||
		strings.Contains(html, "packagequality.com")
}

func isIcon(html string) bool {
	return strings.Contains(html, "emoji") ||
		strings.Contains(html, "icon")
}

func isLogo(html string) bool {
	return strings.Contains(html, "logo")
}
