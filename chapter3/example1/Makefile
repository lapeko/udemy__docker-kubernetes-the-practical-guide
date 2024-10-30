.PHONY: build, build-dev, run, run-dev, clean

build:
	docker build -t feedback:volume .

build-dev:
	docker build -t feedback:volume-dev -f Dockerfile-dev .

run:
	docker run --name feedback --rm -d -v ${PWD}:/app:ro -v feedback:/app/permanent-data -v /app/temporary-data -p 3000:3000 feedback:volume

run-dev:
	docker run --name feedback-dev --rm -d -v ${PWD}:/app:ro -v feedback-dev:/app/permanent-data -v /app/temporary-data -p 3000:3000 feedback:volume-dev

clean:
	docker stop feedback || true && docker rm feedback || true && docker rmi feedback:volume || true && docker volume rm feedback || true

clean-dev:
	docker stop feedback-dev || true && docker rm feedback-dev || true && docker rmi feedback:volume-dev || true && docker volume rm feedback-dev || true
