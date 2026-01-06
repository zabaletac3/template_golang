APP_NAME=api
CMD_DIR=cmd/api
TMP_DIR=tmp
LOG_DIR=logs

AIR_BIN=$(shell go env GOPATH)/bin/air

.PHONY: dev build run clean test lint

## 🔥 Desarrollo con hot reload
dev:
	@echo "🔥 Starting development server with hot reload..."
	@if ! command -v $(AIR_BIN) > /dev/null; then \
		echo "📦 Installing Air..."; \
		go install github.com/cosmtrek/air@v1.49.0; \
	fi
	@mkdir -p $(TMP_DIR) $(LOG_DIR)
	$(AIR_BIN) -c .air.toml

## 🏗️ Build manual
build:
	@echo "🏗️ Building binary..."
	go build -o $(TMP_DIR)/$(APP_NAME) ./$(CMD_DIR)

## ▶️ Run sin hot reload
run: build
	@echo "▶️ Running binary..."
	./$(TMP_DIR)/$(APP_NAME)

## 🧹 Limpieza
clean:
	@echo "🧹 Cleaning artifacts..."
	rm -rf $(TMP_DIR) $(LOG_DIR)

## 🧪 Tests
test:
	go test ./... -race -count=1

## 🔍 Lint (requiere golangci-lint)
lint:
	golangci-lint run
