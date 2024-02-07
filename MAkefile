build:
	@mkdir -p bin
	@go build -o bin/goblock

run: build
	@./bin/docker

.PHONY: test
test:
	@go test -v ./...


