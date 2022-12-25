#!/bin/bash

set -e

CUR_REVISION=$(git rev-parse --short=12 HEAD)

jq -n --arg rev "$CUR_REVISION" '{"revision":$rev}'

set +e
