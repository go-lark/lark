#!/usr/bin/env bash

set -e

cd ./v2
echo "mode: atomic" >coverage.txt
go test -coverprofile=profile.out -covermode=atomic ./...
if [ -f profile.out ]; then
	tail -q -n +2 profile.out >>coverage.txt
	rm profile.out
fi

go tool cover -func coverage.txt
