#!/bin/bash

CONTAINER_NAME="go-users-micro-service"

# stop and remove container if it exists
docker rm $(docker stop $(docker ps -a -q --filter ancestor=$CONTAINER_NAME --format="{{.ID}}"))
# remove image if it exists
docker images -a | grep "$CONTAINER_NAME" | awk '{print $3}' | xargs docker rmi
# create container
docker build -t $CONTAINER_NAME .
# run container
docker run -d -p 8080:8080 $CONTAINER_NAME