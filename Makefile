MONGO_SERVICE = mongodb

swag-v1: ### swag init
	swag init -g internal/controller/http/v1/router.go
.PHONY: swag-v1

mongo-up:
	sudo docker-compose up --build $(MONGO_SERVICE)
.PHONY: mongo-up

mongo-stop:
	sudo docker-compose stop $(MONGO_SERVICE)
.PHONY: mongo-stop

run-app: swag-v1
	go mod tidy && go mod download && \
	DISABLE_SWAGGER_HTTP_HANDLER='' GIN_MODE=debug CGO_ENABLED=0 go run ./cmd/app
.PHONY: run
