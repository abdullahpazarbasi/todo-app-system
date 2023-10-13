#!/usr/bin/env sh

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}" || exit 1

CERT_PATH="${DIRECTORY_PATH_OF_THIS_SCRIPT}/certs/todo.local.crt"

if minikube status --format='{{.Host}}' | grep "Running" 1> /dev/null 2> /dev/null; then
  echo "Minikube is running."
else
  minikube start
  minikube addons enable ingress
  minikube addons enable dashboard
  minikube addons enable metrics-server

  minikube cp $CERT_PATH /tmp/todo.local.crt \
  && minikube ssh "sudo cp -f /tmp/todo.local.crt /usr/share/ca-certificates/todo.local.crt" \
  && minikube ssh "cat /etc/ca-certificates.conf | grep -q \"todo.local.crt\" || echo \"todo.local.crt\" | sudo tee -a /etc/ca-certificates.conf" \
  && minikube ssh "sudo update-ca-certificates" || exit 2

  minikube ssh "sudo mkdir -p /etc/docker/certs.d/registry.todo.local:5000" \
  && minikube ssh "sudo cp -f /tmp/todo.local.crt /etc/docker/certs.d/registry.todo.local:5000/todo.local.crt" \
  && minikube ssh "sudo systemctl restart docker" || exit 3
fi

exit 0
