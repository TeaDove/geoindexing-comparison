GO ?= GO111MODULE=on CGO_ENABLED=1 go

run-backend-manager:
	$(GO) run pkg/entrypoints/manager/main.go

run-backend-worker:
	$(GO) run pkg/entrypoints/worker/main.go


build-backend-manager:
	$(GO) build -o bootstrap pkg/entrypoints/manager/main.go

build-backend-worker:
	$(GO) build -o bootstrap pkg/entrypoints/worker/main.go

show-pprof:
	rm profile001.pdf | true
	$(GO) tool pprof -pdf cpu.prof
	open profile001.pdf
	rm profile001.pdf

test:
	$(GO) test ./...

lint:
	gofumpt -w .
	golines --base-formatter=gofumpt --max-len=120 --no-reformat-tags -w .
	wsl --fix ./...
	golangci-lint run --fix

test-and-lint: test lint