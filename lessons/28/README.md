docker run hello-world
docker --help
docker images

docker search ubuntu
docker pull ubuntu

docker rmi ubuntu:18.04

docker run -it ubuntu bash
docker run -it ubuntu date

docker run --publish 8080:80 nginx (80 - порт nginx, 8080 - порт на хост машине)

docker ps -a