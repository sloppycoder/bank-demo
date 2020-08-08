#!/bin/sh

set -e

echo starting slave with $BOOM_PARAMS
./slave $BOOM_PARAMS &

locust --master --headless --users "${USERS:-1}" 2>&1


