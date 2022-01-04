#!/bin/sh
  
until PGPASSWORD=postgres psql -h "localhost" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done
