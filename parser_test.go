package main

import (
	"strings"
	"testing"
)

func TestParseNightlyPage(t *testing.T) {
	r := strings.NewReader(SampleNightlyBody)
	trending, err := parseNightlyPage(r)
	if err != nil {
		t.Fatalf("failed parsing on HTML in expected format. error: %v", err)
	}

	gotLen := len(trending.First)
	wantLen := 2
	if gotLen != wantLen {
		t.Fatalf("Found %v first timers, want %v", gotLen, wantLen)
	}

	got := trending.First[0]
	wantURL := "https://github.com/user1/repo1"
	if got.URL != wantURL {
		t.Errorf("trending.First[0].URL = %v, want %v", got.URL, wantURL)
	}
	wantName := "user1/repo1"
	if got.Name != wantName {
		t.Errorf("trending.First[0].Name = %v, want %q", got.Name, wantName)
	}
	wantDesc := "A non existing C library."
	if got.Description != wantDesc {
		t.Errorf("trending.First[0].Description = %v, want %q", got.Description, wantDesc)
	}
	wantStars := 168
	if got.Stars != wantStars {
		t.Errorf("trending.First[0].Stars = %v, want %v", got.Stars, wantStars)
	}
	wantLang := "C"
	if got.Language != wantLang {
		t.Errorf("trending.First[0].Language = %v, want %q", got.Language, wantLang)
	}

	got = trending.First[1]
	wantURL = "https://github.com/user2/repo2"
	if got.URL != wantURL {
		t.Errorf("trending.First[1].URL = %v, want %v", got.URL, wantURL)
	}

	gotLen = len(trending.New)
	wantLen = 1
	if gotLen != wantLen {
		t.Errorf("Found %v new repos, want %v", gotLen, gotLen)
	}
	got = trending.New[0]
	wantURL = "https://github.com/user3/repo3"
	if got.URL != wantURL {
		t.Errorf("trending.New[0].URL = %v, want %v", got.URL, wantURL)
	}

	gotLen = len(trending.Repeaters)
	gotLen = 1
	if gotLen != gotLen {
		t.Errorf("Found %v repeat performers, want %v", gotLen, gotLen)
	}
	got = trending.Repeaters[0]
	wantURL = "https://github.com/user4/repo4"
	if got.URL != wantURL {
		t.Errorf("trending.Repeaters[0].URL = %v, want %v", got.URL, wantURL)
	}
}
