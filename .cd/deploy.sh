#!/usr/bin/env sh

RELEASE_NAME="todo-minikube"

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}" || exit 1

eval "$(minikube docker-env)"

helm install $RELEASE_NAME ./todo

exit 0
