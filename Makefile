MONGO_SERVICE = mongodb
NEO4J_SERVICE = neo4j
CLICKHOUSE_SERVICE = clickhouse

swag-v1: ### swag init
	swag init -g internal/controller/http/v1/router.go
.PHONY: swag-v1

mongo-up:
	docker compose -f ./mongo-cluster/docker-compose.yml up
.PHONY: mongo-up

mongo-stop:
	docker compose -f ./mongo-cluster/docker-compose.yml down
.PHONY: mongo-stop

neo4j-up:
	docker compose -f ./neo4j-cluster/docker-compose.yml up
.PHONY: neo4j-up

neo4j-stop:
	docker compose -f ./neo4j-cluster/docker-compose.yml down
.PHONY: neo4j-stop

clickhouse-up:
	docker compose -f ./clickhouse-cluster/docker-compose.yml up
.PHONY: clickhouse-up

clickhouse-stop:
	docker compose -f ./clickhouse-cluster/docker-compose.yml down
.PHONY: clickhouse-stop

run-app: swag-v1
	go mod tidy && go mod download && \
	DISABLE_SWAGGER_HTTP_HANDLER='' GIN_MODE=debug CGO_ENABLED=0 go run ./cmd/app
.PHONY: run-app


stop-app: neo4j-stop clickhouse-stop mongo-stop
