docker_up:
	docker build -t exagonalapp:latest  . 
	docker-compose -f docker-compose.yml -p exgapp up

nerdctl_up:
	nerdctl build -t exagonalapp:latest . 
	nerdctl compose -f docker-compose.yml -p exgapp up

docker_down:
	docker-compose -f docker-compose.yml -p exgapp down --remove-orphans

nerdctl_down:
	nerdctl compose -f docker-compose.yml -p exgapp down --remove-orphans