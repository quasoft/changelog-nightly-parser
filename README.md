# changelog-nightly-parser

[![GoDoc](https://godoc.org/github.com/quasoft/changelog-nightly-parser?status.svg)](https://godoc.org/github.com/quasoft/changelog-nightly-parser) [![Build Status](https://travis-ci.org/quasoft/changelog-nightly-parser.png?branch=master)](https://travis-ci.org/quasoft/changelog-nightly-parser) [![Coverage Status](https://coveralls.io/repos/github/quasoft/changelog-nightly-parser/badge.svg?branch=master)](https://coveralls.io/github/quasoft/changelog-nightly-parser?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/quasoft/changelog-nightly-parser)](https://goreportcard.com/report/github.com/quasoft/changelog-nightly-parser)

A lambda function that parses the [Changelog Nightly](http://nightly.changelog.com/) page for the last night (`http://nightly.changelog.com/YYYY/MM/DD`) and uploads to GitHub the list of trending repositories found (as a JSON file).

# How to use
Define the following environment variables in configuration of AWS Lambda:
- `GITHUB_REPOSITORY` - name of repository to which to upload the JSON file (eg. "trending-daily").
- `GITHUB_OWNER` - Github username (eg. "myusername")
- `GITHUB_TOKEN` - Github personal token (eg. "myusername")

# How to build

First build the application as linux executable:

    GOOS=linux GOARCH=amd64 go build -o main main.go
    zip main.zip main

or


    GOOS=linux GOARCH=amd64 go build -o main main.go
    build-lambda-zip.exe -o main.zip main

if using Windows as build environment.

Then deploy `main.zip` via the AWS console or the cli tool.
