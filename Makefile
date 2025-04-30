lint:
	gofumpt -w *.go
	golines --base-formatter=gofumpt --max-len=120 --no-reformat-tags -w .
	wsl --fix ./...
	golangci-lint run --fix

run-backend:
	go run main.go

show-pprof:
	go tool pprof -web cpu.prof
