mocks:
	GO111MODULE=on mockgen -source=./internal/database/database.go -destination=./internal/test/mocks/database/database_mock.go
 
file_tree:
	tree -I vendor

tests:
	go test -v  ./...

docker_build:
	docker build -t simpleapp . 

nerdctl_build:
	nerdctl build -t simpleapp . 

docker_up:
	docker build -t simpleapp:latest  . 
	docker-compose -f docker-compose.yml -p smlapp up

nerdctl_up:
	nerdctl build -t simpleapp:latest . 
	nerdctl compose -f docker-compose.yml -p smlapp up

docker_down:
	docker-compose -f docker-compose.yml -p smlapp down --remove-orphans

nerdctl_down:
	nerdctl compose -f docker-compose.yml -p smlapp down --remove-orphans