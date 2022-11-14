FROM golang:1.19.3-alpine3.16 AS BUILD
WORKDIR /app
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./workshop.example.yaml ./workshop.yaml
COPY ./Makefile ./Makefile
RUN apk add --update make
RUN make build.bin

FROM golang:1.19.3-alpine3.16 AS RUN
ENV SUBCOMMAND=${SUBCOMMAND:-web start}
ENV HTTP_PORT=${HTTP_PORT:-8080}
WORKDIR /app
COPY --from=BUILD /app/bin /app/bin
CMD exec /app/bin/workshop ${COMMAND}
