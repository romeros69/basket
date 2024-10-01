MONGO_SERVICE = mongodb
NEO4J_SERVICE = neo4j
CLICKHOUSE_SERVICE = clickhouse

swag-v1: ### swag init
	swag init -g internal/controller/http/v1/router.go
.PHONY: swag-v1

mongo-up:
	sudo docker-compose up --build $(MONGO_SERVICE)
.PHONY: mongo-up

mongo-stop:
	sudo docker-compose stop $(MONGO_SERVICE)
.PHONY: mongo-stop

neo4j-up:
	sudo docker-compose up --build $(NEO4J_SERVICE)
.PHONY: neo4j-up

neo4j-stop:
	sudo docker-compose stop $(NEO4J_SERVICE)
.PHONY: neo4j-stop

clickhouse-up:
	sudo docker-compose up --build $(CLICKHOUSE_SERVICE)
.PHONY: clickhouse-up

clickhouse-stop:
	sudo docker-compose stop $(CLICKHOUSE_SERVICE)
.PHONY: clickhouse-stop

run-app: swag-v1
	go mod tidy && go mod download && \
	DISABLE_SWAGGER_HTTP_HANDLER='' GIN_MODE=debug CGO_ENABLED=0 go run ./cmd/app
.PHONY: run-app
