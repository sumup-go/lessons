TAG=test

.PHONY: build
build: ## Build the application and produce a binary called `test`
	go build .

.PHONY: test
test: ## Test our application
	go test ./...

.PHONY: ensure-golangci-lint
ensure-golangci-lint:
	golangci-lint version | grep "1.42"

.PHONY: lint
lint: ensure-golangci-lint ## Lint the application
	golangci-lint run

.PHONY: watch
watch:
	watchexec -r "go run ."

.PHONY: docker-build
docker-build: ## Build a docker image
	docker build -t=${TAG} .


## Help display.
## Pulls comments from beside commands and prints a nicely formatted
## display with the commands and their usage information.

.DEFAULT_GOAL := help

help: ## Prints this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

