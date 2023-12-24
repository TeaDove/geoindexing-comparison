PYTHON ?= .venv/bin/python

test-integration:
	cd internal && go test ./... --run 'TestIntegration_*' -cover

test-unit:
	cd internal && go test ./... --run 'TestUnit_*' -cover

lint:
	cd internal && golangci-lint run ./...
	cd internal && golines --base-formatter=gofmt --max-len=120 --no-reformat-tags -w .

test: test-unit lint test-integration

run:
	cd internal && go run cases/run/main.go

run-and-save:
	cd internal && go run cases/run/main.go > ../.test.json 2>&1

visualisation-run:
	cd visualisation && docker-compose up -d

visualisation-stop:
	cd visualisation && docker-compose down -d


visualisation-install:
	cd visualisation && python3.10 -m venv .venv
	cd visualisation && $(PYTHON) -m pip install poetry
	cd visualisation && $(PYTHON) -m poetry update

jup:
	cd visualisation && $(PYTHON) -m jupyterlab

jup-darwin:
	cd visualisation && $(PYTHON) -m jupyterlab --app-dir=/opt/homebrew/share/jupyter/lab
