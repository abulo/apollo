#!/bin/bash
find ./ -name ".DS_Store" -depth -exec rm {} \;
docker stop $(docker ps -a -q);
docker rm $(docker ps -a -q);
docker-compose stop;
docker-compose down;
docker system prune --force --volumes;