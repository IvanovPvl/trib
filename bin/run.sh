#!/bin/bash

docker build -t trib .
docker run -p 3333:3333 --name trib-prod trib:latest