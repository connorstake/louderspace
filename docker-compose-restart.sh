#!/bin/bash

# Directory containing the docker-compose.yml file
COMPOSE_DIR="./scripts/docker/"

# Navigate to the directory
cd $COMPOSE_DIR || { echo "Directory not found: $COMPOSE_DIR"; exit 1; }


# Stop and remove Docker containers, networks, and volumes
docker-compose down -v

# Start Docker containers in detached mode
docker-compose up -d

# Display status of running containers
docker-compose ps
