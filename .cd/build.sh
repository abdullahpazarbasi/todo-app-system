#!/usr/bin/env sh

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}" || exit 1

sh build-todo-redis.sh
sh build-todo-service-go.sh
sh build-todo-wbff-go.sh
sh build-todo-mbff-go.sh
sh build-todo-wui.sh

exit 0
