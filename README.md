# Continuous Build, Test and Sign your Containers

A simple REST API built in `golang` using Labstack's [Echo](https://https://echo.labstack.com/]), to demonstrate how integrate CI using,

- [Harness CI](https://app.harness.io)
- [Drone CI](https://drone.io)

Optionally as part of the CI the pipeline we can also sign the container image using [cosign](https://sigstore.dev).

## Pre-requisites

- [Docker Desktop](https://docs.docker.com/desktop/)
- [k3D](https://k3d.io/)
- [Drone CI CLI](https://docs.drone.io/cli/install/)
- [ko](https://ko.build)
- [helm](https://helm.sh)
- [cosign](https://docs.sigstore.dev/cosign/installation)

## Using Harness Platform

Register yourself for a Free Tier Harness Account at <https://app.harness.io>.

To configure the Harness CI pipeline for this project you need the following,

- Docker Registry Account Credentials e.g Docker Hub, Quay.io or Harbor
- GitHub Account with a Personal Access Token (PAT) with `admin:repo` and `user` permissions
- Private and Public Key pair to sign the built container image
- Kubernetes Cluster

## Download Sources

Clone the sources and `cd` into it,

```shell
git clone https://github.com/kameshsampath/go-hello-world.git && cd "$(basename "$_" .git)"
export TUTORIAL_HOME="$PWD"
```

## Setup Environment

Create `dontenv` file that we will be using to set/load our environment variables.

```shell
cp "$TUTORIAL_HOME/.env.example" "$TUTORIAL_HOME/.env"
```

Ensure you update `REPLACE ME` in "$TUTORIAL_HOME/.env" as per your settings.

Spin up a local Kubernetes cluster where we will deploy the demo application.

```shell
"$TUTORIAL_HOME/bin/setup.sh"
```

## Signing and Verify Image

Generate private and public key and save them as kubernetes secret `my-image-sigs` in namespace `cosign-system`,

```shell
kubectl create ns cosign-system
cosign generate-key-pair k8s://cosign-system/my-image-sigs
```

Sign and push the image using Drone CI pipelines,

> **IMPORTANT**: We need to make sure that drone is run with same network as the k3s cluster `$K3D_CLUSTER_NAME`, allowing it to have access to the Cluster kubeconfig

```shell
drone exec --env-file=.env --trusted --network="$K3D_CLUSTER_NAME"
```

Verify image signature,

```shell
drone exec --env-file=.env --trusted --pipeline=verify --network="$K3D_CLUSTER_NAME"
```

## Deploy Kubernetes

Let us use D[sigstore](https://github.com/sigstore/policy-controller) Policy Controller to enforce policy that will allow only signed images to be deployed as part of Kubernetes deployments.

```shell
helm repo add sigstore https://sigstore.github.io/helm-charts
helm repo update
```

Deploy `policy-controller`,

```shell
helm upgrade --install policy-controller \
  -n cosign-system \
  sigstore/policy-controller
```

```shell
kubectl create secret generic my-verify-key -n cosign-system \
  --from-file=cosign.pub="$TUTORIAL_HOME/cosign.pub"
```

Create a `ClusterImagePolicy` that will allow only images signed using keys from `my-verify-key` in `cosign-system`,

```shell
cat <<EOF | kubectl apply -f -
apiVersion: policy.sigstore.dev/v1alpha1
kind: ClusterImagePolicy
metadata:
  name: cip-key-secret
  namespace: cosign-system
spec:
  images:
  - glob: "**"
  authorities:
  - key:
      secretRef:
        name: my-verify-key
EOF
```

Let us create a namespace to deploy the application,

```shell
kubectl create ns demo-apps
```

To enforce policy on all applications in this namespace, label the namespace with `policy.sigstore.dev/include=true`

```shell
kubectl label namespace demo-apps policy.sigstore.dev/include=true
```

Try deploying an image which are not signed using the `my-image-sigs`,

```shell
kubectl run --image library/nginx -n demo-apps nginx
```

The deployment should fail with message like,

```text
Error from server (BadRequest): admission webhook "policy.sigstore.dev" denied the request: validation failed: failed policy: cip-key-secret: spec.containers[0].image
index.docker.io/library/nginx@sha256:b8f2383a95879e1ae064940d9a200f67a6c79e710ed82ac42263397367e7cc4e signature key validation failed for authority authority-0 for index.docker.io/library/nginx@sha256:b8f2383a95879e1ae064940d9a200f67a6c79e710ed82ac42263397367e7cc4e: no matching signatures:
```

Let us now deploy `$IMAGE_REGISTRY/$IMAGE_REGISTRY_USERNAME/go-hello-world:$IMAGE_TAG`,

```shell
kubectl run --image "$IMAGE_REGISTRY/$IMAGE_REGISTRY_USERNAME/go-hello-world:$IMAGE_TAG" -n demo-apps hello-world
```

The application should now be created as the image is signed using the keys from `my-image-sigs`.

## Call API

Do a port-forward to application port `8080`,

```shell
kubectl port-forward -n demo-apps hello-world 8080:8080
```

Try calling the API to see a response `Hello World!`.

```shell
curl http://localhost:8080/
```

The command should return `Hello World!`.

## Cleanup

```shell
"$TUTORIAL_HOME/bin/cleanup.sh"
```
