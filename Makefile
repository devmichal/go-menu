build: clear vendor
	go build -o bin/food-app cmd/*.go

swagger:
	swag init -d "./cmd/server,./internal/food/application" -o ./api --ot json

test: vendor
	docker-compose exec app go test -v ./...

clear:
	rm -rf bin/

IMAGE_NAME?=food-api
container: build/package/Dockerfile
	docker build -t $(IMAGE_NAME) -f build/package/Dockerfile .

vendor: go.mod go.sum
	GOPRIVATE="github.com/devmichal/pdf" go mod vendor

run:
	nodemon --exec go run cmd/main.go --signal SIGTERM

net:
	docker network create db-fiber

start:
	docker run --network db-fiber -it --rm -w /Users/michalszczurowski/code/go/fiber -v /Users/michalszczurowski/code/go/fiber:/Users/michalszczurowski/code/go/fiber -p 5005:5000 cosmtrek/air

IMAGE_NAME?=fiber:1.0.0
container: build/package/Dockerfile
	docker build -t $(IMAGE_NAME) -f build/package/Dockerfile .
    docker run --name fiber-app -itd -p 8008:8010 $(IMAGE_NAME)

.DEFAULT_GOAL=build
.PHONY=build swagger test clear container