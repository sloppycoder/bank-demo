#!/bin/sh

set -e

SLAVES=$SLAVES
if [ "$SLAVE" = "" ]; then
	SLAVE=1
fi

echo starting $SLAVES slaves with $BOOM_PARAMS
for i in {1..$SLAVES}; do 
	./slave $BOOM_PARAMS &
	sleep 1
done

locust --master --headless --users "${USERS:-1}" 2>&1


