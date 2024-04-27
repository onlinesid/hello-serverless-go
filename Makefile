.PHONY: build clean deploy

build:
	
	# install dependencies
	go mod tidy

	# build the binary with no debug information
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bootstrap api/main.go

clean:
	rm -rf ./vendor Gopkg.lock bootstrap

deploy: clean build
	sls deploy --verbose
