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
	docker-compose -f docker-compose.yml up

nerdctl_up:
	nerdctl compose up