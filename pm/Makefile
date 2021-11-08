.PHONY: setup
setup:
	go mod download

.PHONY: build
build:
	go build -ldflags="-w -s" -o ./bin/main ./cmd/pm/main.go

.PHONY: run
run:
	GO_ENV="local" LOCAL_SERVER_PORT=8080 LOCAL_CORS_ALLOWED_ORIGINS="http://localhost:3000" LOCAL_MYSQL_DSN="root:root@tcp(127.0.0.1:3336)/test_db?parseTime=true" go run cmd/pm/main.go
