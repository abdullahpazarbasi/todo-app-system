#!/usr/bin/env sh

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}" || exit 1

sh generate-certs.sh && sh start-minikube.sh && sh boot-up-image-registry.sh && sh set-hosts.sh && sh build.sh && sh deploy.sh

exit 0
