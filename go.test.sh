#!/usr/bin/env bash

# From https://github.com/codecov/example-go

set -e
echo "" > coverage.txt

for d in $(find . -maxdepth 10 -type d); do
    if ls $d/*.go &> /dev/null; then
        go test -coverprofile=profile.out -covermode=atomic $d
        if [ -f profile.out ]; then
            cat profile.out >> coverage.txt
            rm profile.out
        fi
    fi
done
