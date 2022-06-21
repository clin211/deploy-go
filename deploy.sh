#! /bin/bash

set -eux

containerName="deploy-go"

if [[ -z $(docker ps -q -f "name=^${containerName}$") ]]; then
    echo "container not exist"
else
    echo "container exist"
    echo "------------stop container------------"
    docker stop ${containerName}
    echo "------------rm container------------"
    docker rm ${containerName}
    echo "------------print container with all------------"
    docker ps -a
    echo "------------remove image------------"
    docker rmi "forest211/"${containerName}":latest"
    docker images
    echo "------------reload container------------"
    docker-compose up -d
fi
