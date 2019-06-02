.PHONY: clean build deploy

clean:
	rm -rf ./bin

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/autoclean-tagger *.go

deploy: clean build
	sls deploy --verbose
