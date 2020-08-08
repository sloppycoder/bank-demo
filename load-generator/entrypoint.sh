#!/bin/sh

set -e

# for some reason slave cannot connect to database
# without this pause
sleep 3

echo starting slave with $BOOM_PARAMS
./slave $BOOM_PARAMS &

locust --master --headless --users "${USERS:-1}" 2>&1


