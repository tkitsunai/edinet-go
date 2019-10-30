.PHONY : buildrun build run push

buildrun: build run

clean:
	-docker stop edinet-go-redis
	-docker rm edinet-go-redis
	-docker rmi tkitsunai/edinet-go-redis:5.0.6

build:
	docker build -t tkitsunai/edinet-go-redis:5.0.6 .

run:
	-docker stop edinet-go-redis
	-docker rm edinet-go-redis
	docker run --name edinet-go-redis -p6379:6379 -d tkitsunai/edinet-go-redis:5.0.6

push:
	-docker login
	docker push tkitsunai/edinet-go-redis:5.0.6
