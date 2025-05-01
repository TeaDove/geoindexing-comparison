GO ?= GO111MODULE=on CGO_ENABLED=1 go

run-backend:
	$(GO) run main.go

show-pprof:
	$(GO) tool pprof -web cpu.prof

test:
	$(GO) test ./...

lint:
	gofumpt -w .
	golines --base-formatter=gofumpt --max-len=120 --no-reformat-tags -w .
	wsl --fix ./...
	golangci-lint run --fix

test-and-lint: test lint