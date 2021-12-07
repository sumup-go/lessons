.PHONY: gen
gen: ## Generate the API from spec
	oapi-codegen -package main spec.yaml > gen.go

.PHONY: run
run: ## Run the binary
	go run .

.PHONY: test-migrate
test-migrate: ## Test the migrations
	# bring up postgres DB
	# run the migrations in the up direction
	# run the migrations in the down direction
	docker run -d \
    --name postgres \
		-p 5432:5432 \
    -e POSTGRES_PASSWORD=postgres \
    postgres
	./pg-wait.sh
	migrate -path=migrations -database=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable up
	docker kill postgres
	docker rm postgres

## Help display.
## Pulls comments from beside commands and prints a nicely formatted
## display with the commands and their usage information.

.DEFAULT_GOAL := help

help: ## Prints this help
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
