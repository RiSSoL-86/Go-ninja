database:
	docker compose -f compose.deps.dev.yml up -d

install-pre-commit:
	pre-commit install

lint:
	go fmt ./...
	go vet ./...

tests:
	go test -race -v ./...

build:
	go build -race -o bin/app.exe ./src

run:
	go run -race ./src
