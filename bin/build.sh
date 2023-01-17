#!/usr/bin/env bash

set -euxo pipefail

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
export IMAGE_DIGESTS_FILE="$SCRIPT_DIR/../image-refs.txt"
export KUBECONFIG="$SCRIPT_DIR/../.kube/config.internal"
k3d kubeconfig get "$K3D_CLUSTER_NAME" > "$KUBECONFIG"

k3d kubeconfig get "${K3D_CLUSTER_NAME}" > "${KUBECONFIG}"
echo "" > "$IMAGE_DIGESTS_FILE"

echo "$IMAGE_REGISTRY_PASSWORD" | ko login $IMAGE_REGISTRY -u "$IMAGE_REGISTRY_USERNAME" --password-stdin
ko build --bare --tags="$IMAGE_TAG" --platform=linux/amd64 --platform=linux/arm64 --image-refs="$IMAGE_DIGESTS_FILE"

while read -r img_ref; do cosign sign --key="$COSIGN_PRIVATE_KEY" "$img_ref" | jq .; done < "$IMAGE_DIGESTS_FILE"
