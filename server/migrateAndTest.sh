#!/bin/sh

echo "Running migrations..."
go run ./consumer/database/migration/run/main.go
if [ $? -ne 0 ]; then
  echo "Migration failed. Exiting..."
  exit 1
fi
sleep 2
echo "Running tests..."
go test ./...
exit $?