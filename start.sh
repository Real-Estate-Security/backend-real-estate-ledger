#!/bin/sh

set -e

echo "run db migrations"
# Run migrations


echo "start app"
exec "$@"