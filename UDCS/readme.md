## User Data Control Server
Ubuntu 20.04 LTS

## Setup Spring
* cd setup
* sh setupSpring.sh

The default password for MySQL is set in the installation script, so please change it yourself for security reasons.

## Setup ELK
* cd setup
* sh setupElk.sh

## Start Spring
Please ensure the ./setup/env/spring.env file is created before starting Spring.
* sh startSpring.sh

## Start ELK
* sh startELK.sh