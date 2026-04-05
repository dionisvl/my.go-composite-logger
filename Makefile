.PHONY: build up down logs test shell clean

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

shell:
	docker compose exec app sh

test:
	docker compose run --rm app go test ./... -v

clean:
	docker compose down --rmi local
