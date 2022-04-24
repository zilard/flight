.PHONY: image clean run start stop

NAME = flight

clean: stop
	docker ps -a | grep ${NAME}-image:latest | awk '{print $$1 }' | xargs -I {} docker rm {} --force | true
	docker rmi ${NAME}-image:latest --force | true
	docker image rm -f `docker images -f dangling=true -q` | true

image:
	docker image build . -t ${NAME}-image:latest

run:
	docker run -d -p 8080:8080 ${NAME}-image:latest

start:      
	docker ps -a | grep ${NAME}-image:latest | awk '{print $$1 }' | xargs -I {} docker start {} | true

stop:
	docker ps -a | grep ${NAME}-image:latest | awk '{print $$1 }' | xargs -I {} docker stop {} | true
