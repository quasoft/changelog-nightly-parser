linux:
	GOOS=linux GOARCH=amd64 go build -o main main.go
	zip main.zip main

win:
	GOOS=linux GOARCH=amd64 go build -o main main.go
	build-lambda-zip -o main.zip main

.PHONY: linux win
