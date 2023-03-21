FROM gcr.io/distroless/static-debian11

COPY server /server

CMD ["/server"]
