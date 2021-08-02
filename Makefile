_docker-compose-build:
	go mod vendor
	docker-compose -f .docker/docker-compose.yaml build

_docker-compose-up:
	docker-compose -f .docker/docker-compose.yaml up --remove-orphans

start-dependencies: _docker-compose-build _docker-compose-up

run: 
	go run .