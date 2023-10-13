#!/usr/bin/env sh

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}/.." || exit 1

cd todo-service-go || exit 2
docker build --tag registry.todo.local:5000/todo-service-go:latest --file Dockerfile .
docker push registry.todo.local:5000/todo-service-go:latest

exit 0
