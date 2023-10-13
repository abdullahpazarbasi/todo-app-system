#!/usr/bin/env sh

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}/.." || exit 1

cd todo-mbff-go || exit 2
docker build --tag registry.todo.local:5000/todo-mbff-go:latest --file Dockerfile .
docker push registry.todo.local:5000/todo-mbff-go:latest

exit 0
