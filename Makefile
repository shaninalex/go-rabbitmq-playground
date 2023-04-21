restart:
	docker compose down -v
	docker compose up -d --build

down:
	docker compose down -v
