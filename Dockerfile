FROM alpine:latest AS server

ARG BIN_PATH="gMountie"


COPY  $BIN_PATH /opt/gmountie/gMountie
ENTRYPOINT ["/opt/gmountie/gMountie", "serve"]


