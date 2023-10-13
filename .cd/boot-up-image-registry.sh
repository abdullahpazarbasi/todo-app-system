#!/usr/bin/env sh

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}" || exit 1

CONTAINER_NAME="secure-registry"

if docker ps -a --format '{{.Names}}' | grep -q "${CONTAINER_NAME}"; then
    if docker ps --format '{{.Names}}' | grep -q "${CONTAINER_NAME}"; then
        echo "Container ${CONTAINER_NAME} is running..."
    else
        docker start ${CONTAINER_NAME}
    fi
else
    docker run --name ${CONTAINER_NAME} --detach --restart=always --publish 5000:5000 --volume ${DIRECTORY_PATH_OF_THIS_SCRIPT}/certs:/certs -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/todo.local.crt -e REGISTRY_HTTP_TLS_KEY=/certs/todo.local.key registry:2
fi

exit 0
