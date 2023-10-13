#!/usr/bin/env sh

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}" || exit 1

sh stop-minikube.sh && sh shut-down-image-registry.sh

exit 0
