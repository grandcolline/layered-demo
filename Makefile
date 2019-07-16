.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: logs
logs:
	docker-compose logs -f

.PHONY: build
build:
	docker build -t grandcolline/layered-demo .

.PHONY: cleanup
cleanup:
	docker-compose down
	docker rmi -f layered-demo_app
	docker rmi -f grandcolline/layered-demo
