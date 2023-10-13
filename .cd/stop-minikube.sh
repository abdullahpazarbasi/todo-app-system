#!/usr/bin/env sh

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}" || exit 1

if minikube status --format='{{.Host}}' | grep "Stopped" 1> /dev/null 2> /dev/null; then
  echo "Minikube is not running!"
else
  sh undeploy.sh

  echo "Minikube is stopping..."
  minikube stop
fi

exit 0
