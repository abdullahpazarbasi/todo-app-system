#!/usr/bin/env sh

HOST_NAMES="www.todo.local web.bff.todo.local mobile.bff.todo.local"

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}" || exit 1

if cat /etc/hosts | grep "# TODO APP SYSTEM" 1> /dev/null 2> /dev/null; then
  echo "The hosts are already set in host"
else
  if minikube status --format='{{.Host}}' | grep "Running" 1> /dev/null 2> /dev/null; then
    HOST_IPvX="$(minikube ip)"
    printf "\n\n# TODO APP SYSTEM\n${HOST_IPvX} ${HOST_NAMES}\n" | sudo tee -a /etc/hosts
  else
    echo "minikube is not running!"
  fi
  HOST_IPvX="$(docker inspect secure-registry --format='{{.NetworkSettings.Gateway}}')"
  printf "${HOST_IPvX} registry.todo.local\n" | sudo tee -a /etc/hosts
fi

if minikube status --format='{{.Host}}' | grep "Running" 1> /dev/null 2> /dev/null; then
  if [ `minikube ssh "cat /etc/hosts | grep -q \"# TODO APP SYSTEM\" && echo \"1\" || echo \"\""` = "1" ]; then
    echo "The hosts are already set in minikube"
  else
    INTERFACE="$(ip route | grep default | awk '{print $5}')"
    HOST_IPvX="$(ip addr show ${INTERFACE} | grep 'inet ' | awk '{print $2}' | cut -d/ -f1)"
    minikube ssh "printf \"\\n\\n# TODO APP SYSTEM\\n${HOST_IPvX} registry.todo.local\\n\" | sudo tee -a /etc/hosts"
  fi
else
  echo "minikube is not running!"
fi

exit 0
