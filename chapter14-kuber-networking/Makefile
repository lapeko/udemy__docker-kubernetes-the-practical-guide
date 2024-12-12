.PHONY: local docker-down

local: docker-down
	docker compose up -d --build && cd ./frontend && npm run dev

docker-down:
	docker compose down -v