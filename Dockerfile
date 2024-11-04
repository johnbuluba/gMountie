FROM alpine:latest AS server

ARG BIN_PATH="gmountie-server"


COPY  $BIN_PATH /opt/gmountie/server
ENTRYPOINT ["/opt/gmountie/server", "serve"]


