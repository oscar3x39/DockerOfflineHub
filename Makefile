.PHONY: up down

up:
	docker-compose up -d --remove-orphans

down:
	docker-compose down --remove-orphans

logs:
	docker-compose logs