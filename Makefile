all: start

.PHONY: build
build:
	docker-compose build app

.PHONY: start
start:
	docker-compose up
