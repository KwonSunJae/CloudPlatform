#!/bin/bash

BASE_PATH=`pwd -P`

DEPLOY_PATH=$BASE_PATH/setup

chmod -v +x $BASE_PATH/setup/setup/entrypoint.sh

if ! nc -z localhost 5601; then
    echo "ELK is not running on port 5601."

    cd setup

    sudo docker-compose -f $DEPLOY_PATH/docker-compose.yml down
    sleep 5

    sudo docker-compose -f $DEPLOY_PATH/docker-compose.yml up setup
    sleep 10

    sudo docker-compose -f $DEPLOY_PATH/docker-compose.yml up -d
    sleep 20
else
    echo "ELK is Running."
fi
