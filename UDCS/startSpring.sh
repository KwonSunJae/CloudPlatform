#!/bin/bash

BASE_PATH=`pwd -P`
BUILD_PATH=$(ls $BASE_PATH/udcs/build/libs/*.jar 2>/dev/null)

chmod -v +x $BASE_PATH/setup/setup/entrypoint.sh

#if [ -z "$BUILD_PATH" ]; then
#    echo "JAR 파일이 존재하지 않습니다. 빌드를 시작합니다..."
#    sudo $BASE_PATH/udcs/gradlew -p $BASE_PATH/udcs bootJar
#    BUILD_PATH=$(ls $BASE_PATH/udcs/build/libs/*.jar)
#fi


JAR_NAME=$(basename $BUILD_PATH)
echo "> build 파일명: $JAR_NAME"

echo "> build 파일 복사"
DEPLOY_PATH=$BASE_PATH/setup
sudo cp $BUILD_PATH $DEPLOY_PATH

echo "> 현재 구동중인 Spring Application pid 확인"
IDLE_PID=$(pgrep -f $JAR_NAME)

if [ -z $IDLE_PID ]
then
  echo "> 현재 구동중인 애플리케이션이 없으므로 종료하지 않습니다."
else
  echo "> kill -15 $IDLE_PID"
  sudo kill -15 $IDLE_PID
  sleep 5
fi

if ! nc -z localhost 8080; then
  echo "> Spring Application을 실행합니다."
  cd setup
  sudo nohup java -jar $DEPLOY_PATH/*.jar &
fi

