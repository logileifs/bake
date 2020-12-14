hello:
	echo "hello world"

build:
	go build -o bake

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bake-linux-amd64

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o bake-darwin-amd64

run:
	go run main.go

release: build-linux build-macos