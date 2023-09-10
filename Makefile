start:
	docker compose\
		-f docker-compose.yml\
		--env-file=.env\
		up -d --build

down:
	docker compose\
		-f docker-compose.yml\
		--env-file=.env\
		down -v