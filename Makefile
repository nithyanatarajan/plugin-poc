.PHONY: help
help: ## Prints help (only for targets with comments)
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

OUT=out

clean:
	rm -rf $(OUT)

build-provider: ## Build add providers
	go build -C providers -o ../$(OUT)/providers

build: build-provider ## Build application
	go build -o ./$(OUT)/driver ./cmd

run: build ## Run application
	./out/driver

lint: ## Lint the application
	golangci-lint run -v --fix
