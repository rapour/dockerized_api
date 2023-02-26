#!/bin/bash

# build the image 
docker build . -t api:latest

# test the created docker image 
docker run --name api01 -e HTTP_PORT=8080 -p 8080:8080 -d api 

# push the image to DockerHub
docker login 
docker tag api:latest rapour/dockerized_api
docker push rapour/dockerized_api
