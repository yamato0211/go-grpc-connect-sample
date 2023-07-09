.PHONY: tools
tools: ## Install required tools.
	echo 'Run go install' && \
	cd ./tools; \
	cat tools.go | grep "_" | awk -F'"' '{print $$2}' | xargs -tI % go install %@latest && \
	cd ../;

.PHONY: run-api
run-api:
	go run ./cmd/api/main.go

.PHONY: run-client
run-client:
	go run ./cmd/client/main.go

.PHONY: evans
evans: ## Run evans.
	evans --proto ./proto/greet.proto --port 8080

.PHONY: go-generate
go-generate:
	go generate ./pkg/...

.PHONY: go-test
go-test:
	go test ./...
