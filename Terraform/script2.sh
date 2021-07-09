#!/bin/bash

apt update
apt upgrade -y
apt install docker.io
docker pull borodun/front
docker pull borodun/back
sudo docker run -p 80:5000 -e API_BASE_URL=http://37.230.196.108:9999/v1/ -d borodun/front
sudo docker run -p 9999:9999 -d borodun/back


