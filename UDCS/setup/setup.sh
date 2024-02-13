echo "시스템 패키지 업데이트 중..."
sudo apt update
echo "시스템 패키지 업데이트 완료."

echo "MySQL 서버 설치 중..."
sudo apt install -y mysql-server
echo "MySQL 서버 설치 완료."

echo "OpenJDK 17 설치 중..."
sudo apt install -y openjdk-17-jdk
echo "OpenJDK 17 설치 완료."

echo "Gradle 설치 중..."
sudo apt install -y gradle
echo "Gradle 설치 완료."

echo "Docker 및 Docker Compose 설치 중..."
sudo apt install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"
sudo apt update
sudo apt install -y docker-ce

sudo curl -L "https://github.com/docker/compose/releases/download/1.28.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
docker-compose --version
echo "Docker 및 Docker Compose 설치 완료."

echo "MySQL root 계정 비밀번호 변경 중..."
sudo mysql -u root -e "ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '1234';"
sudo mysql -u root -e "FLUSH PRIVILEGES;"
echo "MySQL root 계정 비밀번호 변경 완료."
