test:
	go test -count=1 ./...
	rm -rf libgen/libgen
build:
	GOARCH=amd64 GOOS=darwin go build -o libgen-cli-darwin
	GOARCH=amd64 GOOS=linux go build -o libgen-cli-linux
	GOARCH=amd64 GOOS=freebsd go build -o libgen-cli-freebsd
	GOARCH=amd64 GOOS=windows go build -o libgen-cli-windows
bin:
	sudo ln -s libgen-cli /usr/local/bin

.PHONY: test build bin