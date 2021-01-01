#!/usr/bin/env bash

echo "Starting containers..."
docker-compose up -d

sleep 10
echo "Startup ended"
docker ps

echo "Init db"
docker-compose exec roach1 ./cockroach init --insecure
