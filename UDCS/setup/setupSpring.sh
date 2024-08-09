echo "시스템 패키지 업데이트 중..."
sudo apt update
echo "시스템 패키지 업데이트 완료."

echo "MySQL 서버 설치 중..."
sudo apt install -y mysql-server
echo "MySQL 서버 설치 완료."

echo "OpenJDK 17 설치 중..."
sudo apt install -y openjdk-17-jdk
echo "OpenJDK 17 설치 완료."

echo "MySQL root 계정 비밀번호 변경 중..."
sudo mysql -u root -e "ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '1234';"
sudo mysql -u root -e "FLUSH PRIVILEGES;"
echo "MySQL root 계정 비밀번호 변경 완료."

EXTERNAL_IP=$(curl -s ifconfig.me)

ENV_FILE="./env/spring.env"

if [ ! -f "$ENV_FILE" ]; then
    mkdir -p ./env
    cat << EOF > $ENV_FILE
UDCS_SERVER_URL=http://$EXTERNAL_IP:8080
MYSQL_HOST=localhost
MYSQL_NAME=root
MYSQL_PASSWORD=1234
JWT_CLIENT_SECRET=myjwtclientsecret
JWT_ISSUER=myjwtissuer
GO_SERVER_URL=http://117.16.136.172:3000
REFRESH_TOKEN_EXPIRE=360000
ACCESS_TOKEN_EXPIRE=360000
LOGSTASH_DESTINATION=0.0.0.0:50000
EOF
else
    # Update the env/spring.env file.
    sed -i "/^UDCS_SERVER_URL=/c\UDCS_SERVER_URL=http://${EXTERNAL_IP}:8080" $ENV_FILE
    echo "The external IP address has been updated in $ENV_FILE: $EXTERNAL_IP"
fi
