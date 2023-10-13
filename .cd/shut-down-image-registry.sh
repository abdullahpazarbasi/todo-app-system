#!/usr/bin/env sh

CONTAINER_NAME="secure-registry"

if docker ps -a --format '{{.Names}}' | grep -q "${CONTAINER_NAME}"; then
  docker rm -f ${CONTAINER_NAME}
fi

exit 0
