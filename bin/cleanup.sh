#!/usr/bin/env bash

set -euxo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

drone exec --trusted --env-file="$SCRIPT_DIR/../.env" --pipeline=delete-cluster
 