#!/usr/bin/env sh

RELEASE_NAME="todo-minikube"

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}" || exit 1

if minikube status --format='{{.Host}}' | grep "Stopped" 1> /dev/null 2> /dev/null; then
  echo "Minikube is not running!"
else
  eval "$(minikube docker-env)"

  if helm list -q | grep -q "${RELEASE_NAME}" 1> /dev/null 2> /dev/null; then
    echo "Release ${RELEASE_NAME} found. Uninstalling..."
    helm uninstall $RELEASE_NAME
  else
    echo "Release ${RELEASE_NAME} could not be found."
  fi
fi

exit 0
