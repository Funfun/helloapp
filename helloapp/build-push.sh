#!/usr/bin/env bash

set -euo pipefail

main() {
    go mod vendor
    docker build -t helloapp:stable .
    docker tag helloapp:stable quay.io/alexey_medvedchikov/helloapp:latest
    docker push quay.io/alexey_medvedchikov/helloapp:latest
}

main
