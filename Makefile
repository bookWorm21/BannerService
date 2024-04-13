ENV=./config/.env

.PHONY: local-run
local-run:
	go run ./cmd/main.go

.PHONY: go-build
go-build:
	go build -o server ./cmd/main.go

.PHONY: docker-build
docker-build:
	docker build --no-cache --network host -t banner_service_go .

.PHONY: docker-run
docker-run:
	docker compose --env-file $(ENV) up -d

.PHONY: docker-down
docker-down:
	docker compose --env-file $(ENV) down

.PHONY: run-environment
run-environment:
	docker compose --env-file $(ENV) -f ./docker-compose-environment.yml up -d

.PHONY: down-environment
down-environment:
	docker compose --env-file $(ENV) -f ./docker-compose-environment.yml down

.PHONY: docker-clean-data
docker-clean-data:
	docker compose --env-file $(ENV) down
	rm -rf ./.pgdata

.PHONY: local-run-with-env
local-run-with-env:
	make run-environment
	sleep 10 #need wait postgres gets up
	make local-run
