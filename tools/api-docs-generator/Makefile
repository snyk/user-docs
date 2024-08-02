test: tidy lint
	go test -v ./...

run:
	@go run . config.yml ../..

historical-changelog:
	@go run ./cmd/historical-change-log config.yml ../..

lint: tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run

format: tidy
	go run golang.org/x/tools/cmd/goimports -w -local=github.com/snyk/user-docs/tools/api-docs-generator .

tidy:
	go mod tidy

dry-run:
	@go run . config.yml ../.. --dry-run