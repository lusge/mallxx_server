#!/bin/bash

docker run -it --name db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=luxixi1990 -d mysql:8.0