test: tidy lint
	go test ./...

run:
	@go run . config.yml ../..

run-historical-change-log:
	@go run /cmd/historical-change-log

lint: tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint run

format: tidy
	go run golang.org/x/tools/cmd/goimports -w -local=github.com/snyk/user-docs/tools/api-docs-generator .

tidy:
	go mod tidy