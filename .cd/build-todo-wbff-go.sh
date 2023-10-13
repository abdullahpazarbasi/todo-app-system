#!/usr/bin/env sh

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}/.." || exit 1

cd todo-wbff-go || exit 2
docker build --tag registry.todo.local:5000/todo-wbff-go:latest --file Dockerfile .
docker push registry.todo.local:5000/todo-wbff-go:latest

exit 0
