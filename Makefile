# Сборка проекта и запуск тестов всех видов
.PHONY: build test run-ok run-err clean tidy

tidy:
	go mod tidy

build: tidy
	go build -o bin/loglinter ./cmd/app

test:
	go test ./... -v

run-ok: build
	./bin/loglinter ./internal/analyzer/testdata/src/example_ok/...

run-err: build
	./bin/loglinter ./internal/analyzer/testdata/src/example_err/...

lint-with-config: build
	./bin/loglinter -config .loglinter.json ./...

clean:
	rm -rf bin/
