#!/bin/bash

if [ "$(docker ps -a -q -f name=yan-compiler)" ]; then
    echo "Using existing container yan-compiler."
    docker start yan-compiler
else
    echo "Creating and starting new container yan-compiler."
    docker run --name yan-compiler -d compiler
fi