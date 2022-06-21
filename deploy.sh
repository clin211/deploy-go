#!/bin/bash
#
set -eux

containerName="php748"

if [[ -n $(docker ps -q -f "name=^${containerName}$") ]]; then
    echo "container exist"
else
    echo "container not exist"
fi

if [[ -z $(docker ps -q -f "name=^${containerName}$") ]]; then
    echo "container not exist"
else
    echo "container exist"
fi
