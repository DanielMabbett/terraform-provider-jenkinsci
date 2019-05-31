#!/bin/bash -e
echo "[1] Starting the Jenkins Master container:"
docker run -d --rm --name gitlab \
    -e GITLAB_ROOT_PASSWORD=adminadmin \
    -p 8080:8080 -p 50000:50000  \
    jenkins