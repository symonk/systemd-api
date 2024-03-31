FROM golang:1.21-alpine AS build

RUN mkdir -p /go/src/github.com/symonk/systemd-api && \
    apk add --no-cache git openssh-client make gcc libc-dev

WORKDIR /go/src/github.com/symonk/systemd-api
COPY . .
RUN make build

FROM alpine:3
COPY --from=build /go/src/github.com/symonk/systemd-api/systemd-api /bin/systemd-api
CMD /bin/systemdapi
