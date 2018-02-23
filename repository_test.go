package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestRepository_readmeURL(t *testing.T) {
	tests := []struct {
		name string
		r    Repository
		want string
	}{
		{"Short", Repository{URL: "https://github.com/user/repo"}, "https://api.github.com/repos/user/repo/readme"},
		{"With www", Repository{URL: "https://www.github.com/user/repo"}, "https://api.github.com/repos/user/repo/readme"},
		{"Http", Repository{URL: "http://github.com/user/repo"}, "http://api.github.com/repos/user/repo/readme"},
		{"With trailing slash", Repository{URL: "http://www.github.com/user/repo/"}, "http://api.github.com/repos/user/repo/readme"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.readmeURL(); got != tt.want {
				t.Errorf("Repository.readmeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_rawImageURL(t *testing.T) {
	type args struct {
		branch       string
		relativePath string
	}
	tests := []struct {
		name string
		r    Repository
		args args
		want string
	}{
		{
			"Relative image from repository",
			Repository{URL: "https://github.com/user/repo"}, args{branch: "master", relativePath: "images/screenshot.jpg"},
			"https://raw.githubusercontent.com/user/repo/master/images/screenshot.jpg",
		},
		{
			"Relative image from repository with www",
			Repository{URL: "https://www.github.com/user/repo"}, args{branch: "master", relativePath: "images/image.jpg"},
			"https://raw.githubusercontent.com/user/repo/master/images/image.jpg",
		},
		{
			"Relative image from repository with http",
			Repository{URL: "http://github.com/user/repo"}, args{branch: "master", relativePath: "images/demo.png"},
			"http://raw.githubusercontent.com/user/repo/master/images/demo.png",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.rawImageURL(tt.args.branch, tt.args.relativePath); got != tt.want {
				t.Errorf("Repository.rawImageURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_findScreenshot(t *testing.T) {
	downloader = NewStubDownloader()

	tests := []struct {
		name           string
		r              Repository
		httpError      error
		readmeHTML     string
		wantErr        bool
		wantScreenshot string
	}{
		{
			"GET error",
			Repository{URL: "https://github.com/user1/repo1"}, fmt.Errorf("HTTP Error"), ``,
			true, "",
		},
		{
			"No image",
			Repository{URL: "https://github.com/user1/repo1"}, nil, `<p>Just text</p>`,
			true, "",
		},
		{
			"Relative image",
			Repository{URL: "https://github.com/user1/repo1"}, nil, `<img src="screenshot.jpg">`,
			false, "https://raw.githubusercontent.com/user1/repo1/master/screenshot.jpg",
		},
		{
			"Absolute image",
			Repository{URL: "https://github.com/user1/repo1"}, nil, `<img src="http://example.com/demo.jpg">`,
			false, "http://example.com/demo.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			downloader.(*StubDownloader).errorToReturn = tt.httpError
			downloader.(*StubDownloader).body = bytes.NewBufferString(tt.readmeHTML)
			screenshot, err := tt.r.findScreenshot()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.populateScreenshot() error = %v, wantErr %v", err, tt.wantErr)
			} else if screenshot != tt.wantScreenshot {
				t.Errorf("Repository.findScreenshot() = %v, want %v", screenshot, tt.wantScreenshot)
			}
		})
	}
}
