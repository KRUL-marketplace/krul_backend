LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	bin/golangci-lint run ./internal/product-service/... ./internal/user-service/... --config .golangci.pipeline.yaml
