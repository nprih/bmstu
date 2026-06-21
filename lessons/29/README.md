docker compose build
docker compose up -d
docker compose start
docker compose down
docker compose stop
docker compose logs -f [service name]
docker compose ps
docker compose exec [service name] [command]
docker compose images

version: '3.9'
services:
db:
image: mariadb:10.10.2
restart: always
environment:
MYSQL_ROOT_PASSWORD: 12345
adminer:
image: adminer:4.8.1
restart: always
ports:
- 8080:8080