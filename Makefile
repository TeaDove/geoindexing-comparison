lint:
	gofumpt -w *.go
	golines --base-formatter=gofumpt --max-len=120 --no-reformat-tags -w .
	wsl --fix ./...
	golangci-lint run --fix

run:
	go run main.go
