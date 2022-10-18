start:
	docker-compose -f docker/docker-compose.yml up --build -d 

stop:
	docker-compose -f docker/docker-compose.yml down
