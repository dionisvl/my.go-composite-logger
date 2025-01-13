
up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose up --build -d

sh:
	docker compose exec app /bin/sh
