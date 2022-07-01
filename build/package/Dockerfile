FROM golang:1.18.3-alpine3.16 AS build
WORKDIR /build

RUN apk add --no-cache make

COPY . .

RUN make build

FROM alpine:3.16

RUN adduser --disabled-password application \
    && apk add --no-cache ca-certificates

USER application

COPY --from=build --chmod=700 --chown=application /build/bin/food-app /usr/local/bin/food-app

EXPOSE 8008
ENTRYPOINT [ "go", "run", "cmd/*.go" ]
CMD [ "server" ]
