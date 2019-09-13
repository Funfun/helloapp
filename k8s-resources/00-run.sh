#!/usr/bin/env bash

kubectl run helloapp --image quay.io/alexey_medvedchikov/helloapp:latest --dry-run -o yaml > 00-app.yaml
