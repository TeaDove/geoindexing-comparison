test-integration:
	cd internal && go test ./... --run 'TestIntegration_*' -cover

test-unit:
	cd internal && go test ./... --run 'TestUnit_*' -cover

lint:
	cd internal && golangci-lint run

test: test-unit lint test-integration

bench:
	cd internal && go test -benchmem -bench=. -benchtime=100x # | benchstat /dev/stdin

visualisation-run:
	cd visualisation && docker-compose up -d

visualisation-stop:
	cd visualisation && docker-compose down -d
