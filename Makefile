up:
	docker-compose	up	-d

down:
	docker-compose down -v

db:
	docker exec -it db sh

.PHONY: api
api:
	docker exec -it api sh
