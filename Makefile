test:
	go test -v -race ./...
	rm -rf cmd/libgen-cli/libgen
	rm -rf libgen/libgen

build:
	go build -v -o libgen

build-travis:
	go build -o artifacts/libgen-cli-linux .
	GOOS=darwin GOARCH=amd64 go build -o artifacts/libgen-cli-macos .
	GOOS=windows GOARCH=amd64 go build -o artifacts/libgen-cli-windows.exe . 
	GOOS=freebsd GOARCH=amd64 go build -o artifacts/libgen-cli-freebsd .

install:
	go install .

.PHONY: test build build-travis bin
