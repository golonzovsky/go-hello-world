#!/usr/bin/env bash
echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
docker push $DOCKER_USER/go-hello
