#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
export IMAGE_DIGESTS_FILE="$SCRIPT_DIR/../image-refs.txt"
export KUBECONFIG="$SCRIPT_DIR/../.kube/config.internal"

if [ ! -f "$KUBECONFIG" ];
then
  k3d kubeconfig get "$K3D_CLUSTER_NAME" > "$KUBECONFIG"
fi

echo "$IMAGE_REGISTRY_PASSWORD" | ko login "$IMAGE_REGISTRY" -u "$IMAGE_REGISTRY_USERNAME" --password-stdin

while read -r img_ref; do cosign verify --key="$COSIGN_PUBLIC_KEY" "$img_ref" | jq .; done < "$IMAGE_DIGESTS_FILE"

