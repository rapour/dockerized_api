#!/bin/bash

docker compose up -d

# to import data to the database
docker compose -f ./import-compose.yml -d 

# now you can retrieve the data from /view endpoint