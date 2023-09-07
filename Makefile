test:
	cd internal && go test ./... --run -cover

bench:
	cd internal && go test -test.bench='BenchmarkUnit_*' -benchmem ./... -run=^a -benchtime=100x

visualisation-run:
	cd visualisation && docker-compose up -d

visualisation-stop:
	cd visualisation && docker-compose down -d
