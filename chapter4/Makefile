.PHONY: build-be, build-fe, run-mongo, run-be, run-fe

build-be:
	cd ./backend && \
	docker build -t chapter4-be .

build-fe:
	cd ./frontend && \
	docker build -t chapter5-fe .

run-mongo:
	docker run --rm -d --name chapter4-mongo \
		--platform linux/amd64 \
 		-v chapter4-mongo:/data/db \
 		-e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
 		-e MONGO_INITDB_ROOT_PASSWORD=secret \
 		--network chapter4 \
 		mongo:4

run-be:
	cd ./backend && \
	docker run --rm -d --name chapter4-be \
		-v ${PWD}/backend:/app \
		-e MONGO_USER=mongoadmin \
		-e MONGO_PASS=secret \
		--network chapter4 \
		-p 3000:3000 \
		chapter4-be

run-fe:
	cd ./frontend && \
	docker run --rm -d --name chapter5-fe \
		-v ${PWD}/frontend:/app \
		-v /app/node_modules \
		--network chapter4 \
		-p 8000:8000 \
		chapter5-fe