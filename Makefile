start:
	docker-compose -f docker/docker-compose.yml up --build -d 

stop:
	docker-compose -f docker/docker-compose.yml down

tools:
	GOBIN=${PWD}/bin go install github.com/swaggo/swag/cmd/swag@latest

swagger:
	./bin/swag init --dir ./service/cmd/service/ -o ./swagger --pd