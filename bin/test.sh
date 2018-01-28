#!/bin/bash

docker build -t trib-test -f Dockerfile.test .
docker run --rm --name trib-test trib-test:latest