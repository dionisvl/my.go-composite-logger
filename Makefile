.PHONY: tools check vet sec vuln test build run docker-build up down logs sh clean

tools:
	go install github.com/securego/gosec/v2/cmd/gosec@v2.25.0
	GOTOOLCHAIN=go1.26.1 go install golang.org/x/vuln/cmd/govulncheck@v1.1.4

check: vet sec vuln
	@echo "✓ All checks passed"

vet:
	go vet ./...

sec:
	gosec ./...

vuln:
	govulncheck ./...

test:
	go test ./... -v -race -coverprofile=coverage.out

build:
	go build -o logger ./cmd/app

run:
	go run ./cmd/app

docker-build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

sh:
	docker compose exec app sh

clean:
	docker compose down --rmi local
