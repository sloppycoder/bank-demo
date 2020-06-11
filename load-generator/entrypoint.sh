#!/bin/sh

set -e
trap "exit" TERM

if [ -z "${SVC_ADDR}" ]; then
    SVC_ADDR="192.168.39.251:31400"
fi

locust --host="$SVC_ADDR" --headless --users "${USERS:-1}" 2>&1


