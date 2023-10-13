#!/usr/bin/env sh

CERT_PATH="certs/todo.local.crt"
KEY_PATH="certs/todo.local.key"
CNF_PATH="certs/todo.local.cnf"

DIRECTORY_PATH_OF_THIS_SCRIPT="$(dirname "$0")"
cd "${DIRECTORY_PATH_OF_THIS_SCRIPT}" || exit 1

rm -f ${CERT_PATH} ${KEY_PATH}
openssl req \
  -newkey rsa:4096 -nodes -sha256 -keyout ${KEY_PATH} \
  -x509 -days 365 -out ${CERT_PATH} \
  -config ${CNF_PATH}

PLATFORM="$(uname)"

if [ "${PLATFORM}" = "Darwin" ]; then
  echo "Self-signed certificate will be made TRUSTED:"

  sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ${CERT_PATH}
elif [ "${PLATFORM}" = "Linux" ]; then
  echo "Self-signed certificate will be made TRUSTED:"

  sudo cp -f ${CERT_PATH} /usr/share/ca-certificates/todo.local.crt
  cat /etc/ca-certificates.conf | grep -q "todo.local.crt" || echo "todo.local.crt" | sudo tee -a /etc/ca-certificates.conf
  sudo update-ca-certificates --verbose | grep "todo"

  sudo mkdir -p /etc/docker/certs.d/registry.todo.local:5000
  sudo cp -f ${CERT_PATH} /etc/docker/certs.d/registry.todo.local:5000/
  sudo service docker restart
else
  echo "Unsupported platform: ${PLATFORM}"
  exit 2
fi

exit 0
