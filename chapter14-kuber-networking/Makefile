.PHONY: local docker-down build push kub-del kub-apply run service

local: docker-down
	docker compose up -d --build && cd ./frontend && npm run dev

docker-down:
	docker compose down -v

build:
	cd ./frontend && docker build -t lapeko/frontend .

push:
	docker push lapeko/frontend

kub-del:
	kubectl delete all --all

kub-apply:
	kubectl apply -f ./kubernetes

service:
	minikube service frontend-service

run: build push kub-del kub-apply