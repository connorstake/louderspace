#!/bin/bash

# Directory containing the docker-compose.yml file
COMPOSE_DIR="./scripts/docker/"

# Directory containing the seed.go file
SEED_DIR="../sql/"

COMPOSE_DIR_BACK="../docker/"
# Function to verify and change directory
change_dir() {
  if [ -d "$1" ]; then
    cd "$1" || { echo "Failed to change directory to $1"; exit 1; }
  else
    echo "Directory not found: $1"
    exit 1
  fi
}

# Navigate to the compose directory
change_dir "$COMPOSE_DIR"

# Stop and remove Docker containers, networks, and volumes
docker-compose down -v

# Start Docker containers in detached mode
docker-compose up -d

# Wait for a few seconds to ensure the database is up and running
echo "Waiting for the database to start..."
sleep 10  # Adjust the sleep duration as needed

# Navigate to the seed script directory
change_dir "$SEED_DIR"

# Run the seed script
echo "Running the seed script from $(pwd)..."
go run seed.go

if [ $? -eq 0 ]; then
  echo "Database seeding completed successfully."
else
  echo "Database seeding failed."
  exit 1
fi

# Display status of running containers
cd "$COMPOSE_DIR_BACK"
docker-compose ps