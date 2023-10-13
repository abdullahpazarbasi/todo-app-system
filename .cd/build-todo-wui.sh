#!/usr/bin/env sh

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}/.." || exit 1

cd todo-wui || exit 2
docker build --tag registry.todo.local:5000/todo-wui:latest --target server --file Dockerfile .
docker push registry.todo.local:5000/todo-wui:latest

exit 0
