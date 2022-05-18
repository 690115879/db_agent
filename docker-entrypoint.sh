#!/bin/sh

export DB_PATH=/data/data.db

echo "$TZ" >  /etc/timezone
cp /usr/share/zoneinfo/$TZ   /etc/localtime
exec "$@"