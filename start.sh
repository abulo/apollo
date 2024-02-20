#!/bin/bash
find ./ -name ".DS_Store" -depth -exec rm {} \;
docker system prune --force --volumes;
docker-compose up -d;
