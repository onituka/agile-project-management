RED=\033[31m
GREEN=\033[32m
RESET=\033[0m

COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

.PHONY: setup
setup:
	@go mod download

.PHONY: build
build:
	@go build -ldflags="-w -s" -o ./bin/main ./cmd/pm/main.go

.PHONY: run
run:
	@GO_ENV="local" LOCAL_SERVER_PORT=8080 LOCAL_CORS_ALLOWED_ORIGINS="http://localhost:3000" LOCAL_MYSQL_DSN="root:root@tcp(127.0.0.1:3336)/test_db?parseTime=true" \
	go run cmd/pm/main.go

.PHONY: lint
lint:
	@golangci-lint run --path-prefix=pm --config ../.github/.golangci.yaml ./...

test:
	@GO_ENV="local" LOCAL_SERVER_PORT=8080 LOCAL_CORS_ALLOWED_ORIGINS="http://localhost:3000" LOCAL_MYSQL_DSN="root:root@tcp(127.0.0.1:3336)/test_db?parseTime=true" \
	go test -race -v ./... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

# goldenテストデータを更新
test-update:
	@GO_ENV="local" LOCAL_SERVER_PORT=8080 LOCAL_CORS_ALLOWED_ORIGINS="http://localhost:3000" LOCAL_MYSQL_DSN="root:root@tcp(127.0.0.1:3336)/test_db?parseTime=true" \
	go test ./interfaces/handler/... -update

# goldenテストデータを削除して新規作成
test-clean:
	@GO_ENV="local" LOCAL_SERVER_PORT=8080 LOCAL_CORS_ALLOWED_ORIGINS="http://localhost:3000" LOCAL_MYSQL_DSN="root:root@tcp(127.0.0.1:3336)/test_db?parseTime=true" \
	go test ./interfaces/handler/... -update -clean
