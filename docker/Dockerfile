FROM golang:1.19.3-alpine3.16 AS BUILD
WORKDIR /app
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./workshop.example.yaml ./workshop.yaml
COPY ./Makefile ./Makefile
RUN apk add --update make
RUN make build.server

FROM golang:1.19.3-alpine3.16 AS RUN
ENV EXEC_BIN=${EXEC_BIN:-server}
ENV HTTP_PORT=${HTTP_PORT:-8080}
WORKDIR /app
COPY --from=BUILD /app/server /app/server
CMD exec /app/${EXEC_BIN}
