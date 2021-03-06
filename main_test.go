package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	SampleNightlyBody = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
  <head>
    <title>Changelog Nightly - 2018-02-08</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
  </head>
  <body>   
      <table id="top-all-firsts" class="wrapper" width="100%" cellpadding="0" cellspacing="0" border="0">
        <tr>
          <td width="100%">
            <table width="540" cellpadding="20" cellspacing="0" border="0" align="center">
              <tr>
                <td class="section" width="540">
                  <h2>Top Starred Repositories &ndash; First Timers</h2>

                  <p>These repos were not previously featured in Changelog Nightly</p>

                  <div class="repositories">
                  
                    <div class="repository ">
                      <table>
  <tr class="stats">
    <td width="32" valign="top">
      <a href="https://github.com/user1" title="View user1 on GitHub">
        <img class="avatar" src="https://avatars3.githubusercontent.com/u/1234?v=4" width="20" height="20">
      </a>
    </td>
    <td valign="middle">
      <p>
        <span title="Total Stars"><img height="10" alt="Star" src="/images/star.png" />&nbsp;168</span>
        &nbsp;&nbsp;
        <span title="New Stars"><img height="10" alt="Up" src="/images/up.png" />&nbsp;90</span>
      
      
        &nbsp;&nbsp;
        <span title="Language"><a class="repository-language c" href="https://github.com/trending/c" title="View other trending C repos on GitHub"><span class="dot"></span>C</a></span>
      
      </p>
    </td>
  </tr>
  <tr class="about">
    <td width="32" valign="top">
    </td>
    <td valign="top">
      <h3>
        <a href="https://github.com/user1/repo1" title="View REPO1 on GitHub">user1/repo1</a>
      </h3>
      <p>
        A non existing C library.
      </p>
    </td>
  </tr>
</table>

                    </div>

                    <div class="repository last-of-type">
                      <table>
  <tr class="stats">
    <td width="32" valign="top">
      <a href="https://github.com/user2" title="View user2 on GitHub">
        <img class="avatar" src="https://avatars1.githubusercontent.com/u/2345?v=4" width="20" height="20">
      </a>
    </td>
    <td valign="middle">
      <p>
        <span title="Total Stars"><img height="10" alt="Star" src="/images/star.png" />&nbsp;49</span>
        &nbsp;&nbsp;
        <span title="New Stars"><img height="10" alt="Up" src="/images/up.png" />&nbsp;97</span>
      
      
        &nbsp;&nbsp;
        <span title="Language"><a class="repository-language css" href="https://github.com/trending/css" title="View other trending CSS repos on GitHub"><span class="dot"></span>CSS</a></span>
      
      </p>
    </td>
  </tr>
  <tr class="about">
    <td width="32" valign="top">
    </td>
    <td valign="top">
      <h3>
        <a href="https://github.com/user2/repo2" title="View REPO2 on GitHub">user2/repo2</a>
      </h3>
      <p>
        Next to non existing repo 2.
      </p>
    </td>
  </tr>
</table>

                    </div>
                  
                  </div>
                </td>
              </tr>
            </table>
          </td>
        </tr>
      </table>
  
      <table id="top-new" class="wrapper" width="100%" cellpadding="0" cellspacing="0" border="0">
        <tr>
          <td width="100%">
            <table width="540" cellpadding="20" cellspacing="0" border="0" align="center">
              <tr>
                <td class="section" width="540">
                  <h2>Top New Repositories</h2>

                  <p>These repos were open sourced on February 08, 2018</p>

                  <div class="repositories">
                  
                    <div class="repository ">
                      <table>
  <tr class="stats">
    <td width="32" valign="top">
      <a href="https://github.com/user3" title="View user3 on GitHub">
        <img class="avatar" src="https://avatars1.githubusercontent.com/u/3456?v=4" width="20" height="20">
      </a>
    </td>
    <td valign="middle">
      <p>
        <span title="Total Stars"><img height="10" alt="Star" src="/images/star.png" />&nbsp;49</span>
        &nbsp;&nbsp;
        <span title="New Stars"><img height="10" alt="Up" src="/images/up.png" />&nbsp;29</span>
      
      
        &nbsp;&nbsp;
        <span title="Language"><a class="repository-language css" href="https://github.com/trending/css" title="View other trending CSS repos on GitHub"><span class="dot"></span>CSS</a></span>
      
      </p>
    </td>
  </tr>
  <tr class="about">
    <td width="32" valign="top">
    </td>
    <td valign="top">
      <h3>
        <a href="https://github.com/user3/repo3" title="View REPO3 on GitHub">user3/repo3</a>
      </h3>
      <p>
        Three of nothing is better than nothing.
      </p>
    </td>
  </tr>
</table>

                    </div>
                  </div>
                </td>
              </tr>
            </table>
          </td>
        </tr>
      </table>
   
      <table id="top-all-repeats" class="wrapper" width="100%" cellpadding="0" cellspacing="0" border="0">
        <tr>
          <td width="100%">
            <table width="540" cellpadding="20" cellspacing="0" border="0" align="center">
              <tr>
                <td class="section" width="540">
                  <h2>Top Starred Repositories &ndash; Repeat Performers</h2>

                  <p>These repos were previously featured in Changelog Nightly</p>

                  <div class="repositories">
                  
                    <div class="repository ">
                      <table>
  <tr class="stats">
    <td width="32" valign="top">
      <a href="https://github.com/user4" title="View user4 on GitHub">
        <img class="avatar" src="https://avatars2.githubusercontent.com/u/4567?v=4" width="20" height="20">
      </a>
    </td>
    <td valign="middle">
      <p>
        <span title="Total Stars"><img height="10" alt="Star" src="/images/star.png" />&nbsp;265</span>
        &nbsp;&nbsp;
        <span title="New Stars"><img height="10" alt="Up" src="/images/up.png" />&nbsp;377</span>
      
        &nbsp;&nbsp;
        <span title="Times Listed"><img height="10" alt="Eyes" src="/images/eye.png" />&nbsp;3</span>
      
      
      </p>
    </td>
  </tr>
  <tr class="about">
    <td width="32" valign="top">
    </td>
    <td valign="top">
      <h3>
        <a href="https://github.com/user4/repo4" title="View REPO4 on GitHub">user4/repo4</a>
      </h3>
      <p>
        4R - the fourth sample repository.
      </p>
    </td>
  </tr>
</table>

										</div>

										<div class="repository ">
										Should be ignored
										</div>

                  </div>
                </td>
              </tr>
            </table>
          </td>
        </tr>
      </table>
  </body>
</html>`
	SampleReadmeHTML = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html>
		<head>
			<title>Sample Readme</title>
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		</head>
		<body>
			<img src="images/screenshot.jpg" alt="Screenshot">
		</body>
	</html>`
)

// StubDownloader is a stub implementation of the Downloader interface,
// that creates an artificial response with status code 200 and content
// equal to SampleNightlyBody
type StubDownloader struct {
	body               *bytes.Buffer
	statusCodeToReturn int
	errorToReturn      error
}

func NewStubDownloader() *StubDownloader {
	return &StubDownloader{
		body:               nil,
		statusCodeToReturn: http.StatusOK,
		errorToReturn:      nil,
	}
}

func (s *StubDownloader) Get(url string) (*http.Response, error) {
	body := s.body
	if body == nil {
		body = bytes.NewBufferString(SampleNightlyBody)
	}

	return &http.Response{
		Status:     strconv.Itoa(s.statusCodeToReturn),
		StatusCode: s.statusCodeToReturn,
		Body:       ioutil.NopCloser(body),
		Header:     http.Header{},
	}, s.errorToReturn
}

func (s *StubDownloader) Do(r *http.Request) (*http.Response, error) {
	body := s.body
	if body == nil {
		if strings.Contains(r.URL.Path, "readme") {
			body = bytes.NewBufferString(SampleReadmeHTML)
		} else {
			body = bytes.NewBufferString(SampleNightlyBody)
		}
	}

	return &http.Response{
		Status:     strconv.Itoa(s.statusCodeToReturn),
		StatusCode: s.statusCodeToReturn,
		Body:       ioutil.NopCloser(body),
		Header:     http.Header{},
	}, s.errorToReturn
}

// StubUploader is a stub implementation of the Downloader interface,
// that creates an artificial response with status code 200 and content
// equal to SampleNightlyBody
type StubUploader struct {
	body               *bytes.Buffer
	statusCodeToReturn int
	errorToReturn      error
}

func NewStubUploader() *StubUploader {
	return &StubUploader{
		body:               bytes.NewBufferString(""),
		statusCodeToReturn: http.StatusCreated,
		errorToReturn:      nil,
	}
}

func (s *StubUploader) Do(r *http.Request) (*http.Response, error) {
	io.Copy(s.body, r.Body)

	return &http.Response{
		Status:     strconv.Itoa(s.statusCodeToReturn),
		StatusCode: s.statusCodeToReturn,
		Body:       ioutil.NopCloser(strings.NewReader("")),
		Header:     http.Header{},
	}, s.errorToReturn
}

func TestHandler_NoEnvVariables(t *testing.T) {
	downloader = NewStubDownloader()
	uploader = NewStubUploader()

	// Execute lambda function
	os.Setenv("GITHUB_OWNER", "")
	os.Setenv("GITHUB_REPOSITORY", "trending-daily")
	os.Setenv("GITHUB_TOKEN", "123")
	err := Handler()
	if err == nil {
		t.Fatalf("Should have returned an error when GITHUB_OWNER environment variable is not set.")
	}

	os.Setenv("GITHUB_OWNER", "user")
	os.Setenv("GITHUB_REPOSITORY", "")
	os.Setenv("GITHUB_TOKEN", "123")
	err = Handler()
	if err == nil {
		t.Fatalf("Should have returned an error when GITHUB_REPOSITORY environment variable is not set.")
	}

	os.Setenv("GITHUB_OWNER", "user")
	os.Setenv("GITHUB_REPOSITORY", "trending-daily")
	os.Setenv("GITHUB_TOKEN", "")
	err = Handler()
	if err == nil {
		t.Fatalf("Should have returned an error when GITHUB_TOKEN environment variable is not set.")
	}
}

func TestHandler_DownloadFail(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", "123")
	os.Setenv("GITHUB_OWNER", "user")
	os.Setenv("GITHUB_REPOSITORY", "trending-daily")

	uploader = NewStubUploader()

	downloader = NewStubDownloader()
	downloader.(*StubDownloader).errorToReturn = fmt.Errorf("unexpected error")
	err := Handler()
	if err == nil {
		t.Fatalf("Should have returned an error when download call fails.")
	}
}

func TestHandler_UploadFail(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", "123")
	os.Setenv("GITHUB_OWNER", "user")
	os.Setenv("GITHUB_REPOSITORY", "trending-daily")

	downloader = NewStubDownloader()

	uploader = NewStubUploader()
	uploader.(*StubUploader).errorToReturn = fmt.Errorf("unexpected error")
	err := Handler()
	if err == nil {
		t.Fatalf("Should have returned an error when upload call fails.")
	}

	uploader = NewStubUploader()
	uploader.(*StubUploader).statusCodeToReturn = http.StatusBadRequest
	err = Handler()
	if err == nil {
		t.Fatalf("Should have returned an error when upload request replies with an error.")
	}
}

func TestHandler_OK(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", "123")
	os.Setenv("GITHUB_OWNER", "user")
	os.Setenv("GITHUB_REPOSITORY", "trending-daily")

	downloader = NewStubDownloader()
	uploader = NewStubUploader()

	// Execute lambda function
	err := Handler()
	if err != nil {
		t.Fatalf("failed executing Handler in test. error: %v", err)
	}

	body := uploader.(*StubUploader).body

	// Make sure the body contains a 'contents' field
	got := body.String()
	want := "content"
	if !strings.Contains(got, want) {
		t.Errorf("The PUT body has no 'content' field, body: %s", got)
	}

	// Make sure the body contains a 'committer' field
	want = "committer"
	if !strings.Contains(got, want) {
		t.Errorf("The PUT body has no 'committer' field, body: %s", got)
	}

	// Make sure the body contains a 'committer/name' field
	want = "name"
	if !strings.Contains(got, want) {
		t.Errorf("The PUT body has no 'name' field, body: %s", got)
	}

	// Make sure the body contains a 'committer/email' field
	want = "email"
	if !strings.Contains(got, want) {
		t.Errorf("The PUT body has no 'email' field, body: %s", got)
	}

	// Decode 'content' field from base64
	params := struct {
		Content string `json:"content"`
	}{}
	err = json.Unmarshal(body.Bytes(), &params)
	if err != nil {
		t.Errorf("Decoding PUT body from JSON failed, body: %s", body.String())
	}

	// Make sure the 'content' field contains at least one of the expected repositories
	content, err := base64.StdEncoding.DecodeString(params.Content)
	if err != nil {
		t.Errorf("Decoding 'content' from base64 failed, content: %s", params.Content)
	}
	want = "https://github.com/user1/repo1"
	if !strings.Contains(string(content), want) {
		t.Errorf("The file uploaded does not contain URL '%s', file: %s", want, content)
	}
	// Make sure the 'content' field contains the expected screenshot
	want = "images/screenshot.jpg"
	if !strings.Contains(string(content), want) {
		t.Errorf("The file uploaded does not contain the expected screenshot URL: '%s', file: %s", want, content)
	}
}
