.PHONY: up down restart logs

restart:
	docker-compose restart

up:
	docker-compose up -d --remove-orphans

down:
	docker-compose down --remove-orphans

logs:
	docker-compose logs