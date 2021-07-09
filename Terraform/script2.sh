#!/bin/bash

#apt update
#apt upgrade -y
apt install -y docker.io
docker pull borodun/front
docker pull borodun/back
docker run -p 80:80 -e API_BASE_URL=http://37.230.196.108:9999/v1/ -d borodun/front
docker run -p 9999:9999 -d borodun/back


