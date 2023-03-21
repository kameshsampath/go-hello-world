FROM alpine as dist

WORKDIR /build

ARG TARGETARCH

COPY dist/ /build/

RUN <<EOT
  if [ "$TARGETARCH" = "arm64" ] || [ "$TARGETARCH" = "aarch64" ];
  then
    cp -v go-hello-world_linux_arm64/go-hello-world /build/server_$TARGETARCH
  elif  [ "$TARGETARCH" = "amd64" ] || [ "$TARGETARCH" = "x86_64" ];
  then
    cp -v go-hello-world_linux_amd64_v1/go-hello-world /build/server_$TARGETARCH
  fi
EOT

FROM gcr.io/distroless/static-debian11

ARG TARGETARCH

COPY --from=dist /build/server_$TARGETARCH /server

CMD ["/server"]
