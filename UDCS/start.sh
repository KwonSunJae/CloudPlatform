#!/bin/bash

BASE_PATH=/home/ubuntu/UDCS
BUILD_PATH=$(ls $BASE_PATH/udcs/build/libs/*.jar)
JAR_NAME=$(basename $BUILD_PATH)
echo "> build 파일명: $JAR_NAME"

echo "> build 파일 복사"
DEPLOY_PATH=$BASE_PATH/setup/
sudo cp $BUILD_PATH $DEPLOY_PATH



echo "> 현재 구동중인 Spring Application pid 확인"
IDLE_PID=$(pgrep -f $JAR_NAME)

if [ -z $IDLE_PID ]
then
  echo "> 현재 구동중인 애플리케이션이 없으므로 종료하지 않습니다."
else
  echo "> kill -15 $IDLE_PID"
  kill -15 $IDLE_PID
  sleep 5
fi

if ! nc -z localhost 5601; then
    echo "ELK is not running on port 5601."
    
    cd docker-elk

    sudo docker-compose down
    sleep 5
    
    sudo docker compose up setup --force-recreate
    sleep 10
    
    sudo docker-compose up -d
    sleep 20
else
    echo "Running."
fi

if ! nc -z localhost 8080; then
  echo "> Spring Application을 실행합니다."
  sudo nohup java -jar $DEPLOY_PATH/*.jar &
fi

