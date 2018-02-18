linux:
	GOOS=linux GOARCH=amd64 go build -o main
	zip main.zip main

win:
	GOOS=linux GOARCH=amd64 go build -o main
	build-lambda-zip -o main.zip main

.PHONY: linux win
