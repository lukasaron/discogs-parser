#!/bin/bash

#Example of usage
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=discogs
export DB_SSL_MODE=disable
export DB_USERNAME=user
export DB_PASSWORD=password
export FILE_NAME=./source/releases.xml
export FILE_TYPE=releases
export BLOCK_SIZE=1000
export BLOCK_SKIP=0
export BLOCK_LIMIT=2147483647

go build
./discogs

# TODO
# Filter by quality