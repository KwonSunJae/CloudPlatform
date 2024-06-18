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

echo "MySQL root 계정 비밀번호 변경 중..."
sudo mysql -u root -e "ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '1234';"
sudo mysql -u root -e "FLUSH PRIVILEGES;"
echo "MySQL root 계정 비밀번호 변경 완료."
