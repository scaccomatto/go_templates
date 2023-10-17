docker_up:
	docker build -t hexagonalapp:latest  . 
	docker-compose -f docker-compose.yml -p hexgapp up

nerdctl_up:
	nerdctl build -t hexagonalapp:latest . 
	nerdctl compose -f docker-compose.yml -p hexgapp up

docker_down:
	docker-compose -f docker-compose.yml -p hexgapp down --remove-orphans

nerdctl_down:
	nerdctl compose -f docker-compose.yml -p hexgapp down --remove-orphans

file_tree:
	tree -I vendor