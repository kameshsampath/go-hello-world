#!/usr/bin/env bash

set -euxo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

docker network create --opt "com.docker.network.driver.mtu=1450"\
   "${K3D_CLUSTER_NAME}" || true

drone exec --trusted --env-file="${SCRIPT_DIR}/../.env" --network="${K3D_CLUSTER_NAME}" --pipeline=setup

KUBECONFIG="${SCRIPT_DIR}/../.kube/config"
mkdir -p "$(dirname "$KUBECONFIG")"

k3d kubeconfig get "${K3D_CLUSTER_NAME}" > "${KUBECONFIG}"
sed -i 's|host.docker.internal|127.0.0.1|' "${KUBECONFIG}"