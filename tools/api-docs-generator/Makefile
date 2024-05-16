test: lint
	go test ./...

run:
	go run . config.yml ../..

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run

format:
	go run golang.org/x/tools/cmd/goimports -w -local=github.com/snyk/user-docs/tools/api-docs-generator .
