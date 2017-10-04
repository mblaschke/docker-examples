#!/bin/bash

set -o pipefail  ## trace ERR through pipes
set -o errtrace  ## trace ERR through 'time command' and other functions
set -o nounset   ## set -u : exit the script if you try to use an uninitialised variable
set -o errexit   ## set -e : exit the script if any statement returns a non-true return value

EXAMPLE="$1"
DOCKER_OPTS=""
DOCKER_IMAGE="mblaschke/docker-example:$EXAMPLE"

if [[ ! -d "${EXAMPLE}" ]]; then
    echo "[ERROR] Example not defined"
    exit 1
fi

case "$EXAMPLE" in
	"http-server")
		DOCKER_OPTS="-p 8000:8000"
		;;

	"php")
		DOCKER_OPTS="-p 8000:80"
		;;

	"go-func")
	    docker build -t "mblaschke/docker-example:go-func-runner" go-func/go-funcrunner-base
		DOCKER_OPTS="-p 8000:8000"
		;;
esac

docker build -t "$DOCKER_IMAGE" "$EXAMPLE"
docker run --rm $DOCKER_OPTS "$DOCKER_IMAGE"
