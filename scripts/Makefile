.PHONY: init clean build run

init: 
	go mod download github.com/chromedp/chromedp
	go mod tidy

clean: 
	rm -rf ./bin go.sum
	
build:
	export GO111MODULE=on
	go build -o bin/main cmd/main.go 

run: clean init build
