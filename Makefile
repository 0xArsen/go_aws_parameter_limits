.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/parameterStore parameterStore/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/secretsManager secretsManager/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
