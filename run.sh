#!/bin/bash

# Run the application and build the docker image
echo "Running the application..."
echo ""

docker-compose up -d --build --force-recreate

echo ""
echo "Application is running"