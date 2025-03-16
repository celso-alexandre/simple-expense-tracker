#!/bin/bash
set -e

rm -r query || echo 'Nothing to clean'
export DATABASE_URL=$(cat .env | grep DIRECT_URL | cut -d '=' -f2) 
# echo "DATABASE_URL: $DATABASE_URL"
sqlc generate || (go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest && echo 'Try again')
