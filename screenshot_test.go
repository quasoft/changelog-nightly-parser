package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func Test_screenshotFromHTML(t *testing.T) {
	tests := []struct {
		name string
		html string
		want string
	}{
		{
			"No image",
			`<p>A repository without images</p>`,
			"",
		},
		{
			"Only logo",
			`<img src="images/some-logo.png">`,
			"",
		},
		{
			"Only badges and icons",
			`<img src="images/badge.jpg"><img src="images/icon.jpg">`,
			"",
		},
		{
			"Single image",
			`<img src="images/screenshot.jpg">`,
			"images/screenshot.jpg",
		},
		{
			"Badge and screenshot",
			`<img src="images/badge.jpg"><img src="images/screenshot.jpg">`,
			"images/screenshot.jpg",
		},
		{
			"Emoji and screenshot",
			`<img src="https://assets-cdn.github.com/images/icons/emoji/unicode/1f603.png"><img src="images/screenshot.jpg">`,
			"images/screenshot.jpg",
		},
		{
			"Demo",
			`<img src="images/demo.jpg">`,
			"images/demo.jpg",
		},
		{
			"Sample",
			`<img src="images/some-sample.jpg">`,
			"images/some-sample.jpg",
		},
		{
			"Example",
			`<img src="images/some-example.jpg">`,
			"images/some-example.jpg",
		},
		{
			"Image with absolute URL",
			`<img src="http://example.com/images/demo3.jpg">`,
			"http://example.com/images/demo3.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, err := html.Parse(strings.NewReader(tt.html))
			if err != nil {
				t.Errorf("screenshotFromHTML() failed with error %v", err)
			}

			if got := screenshotFromHTML(root); got != tt.want {
				t.Errorf("screenshotFromHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
