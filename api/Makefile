update-deps:
	go get -u ./...
	go get -u tool
	go mod tidy

test:
	go test -count=1 ./...

build:
	go build -o build/api -a -ldflags="-extldflags '-static'" ./cmd/margincms
