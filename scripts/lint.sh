#!/usr/bin/env bash

rm lint-error.log
touch lint-error.log

set -e

ls -1d ./*/ | xargs golint >> lint-error.log
find ./ -name "*.go" -maxdepth 1 | xargs golint >> lint-error.log
if [ -s ./lint-error.log ]; then
    cat ./lint-error.log
    exit 1
fi
