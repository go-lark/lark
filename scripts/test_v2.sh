#!/usr/bin/env bash

set -e
echo "mode: atomic" >coverage.txt

cd ./v2
go test -coverprofile=profile.out -covermode=atomic ./...
if [ -f profile.out ]; then
	tail -q -n +2 profile.out >>coverage.txt
	rm profile.out
fi

go tool cover -func coverage.txt
