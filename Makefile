.PHONY: test
test:
	golangci-lint run
	go test -v ./...
