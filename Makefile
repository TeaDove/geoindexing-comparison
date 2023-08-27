test:
	cd internal && go test ./... --run -cover

visualisation-run:
	cd visualisation && docker-compose up -d

visualisation-stop:
	cd visualisation && docker-compose down -d
