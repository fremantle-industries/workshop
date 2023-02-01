FROM golang:1.18.10-alpine3.17 AS BUILD_GOLANG
WORKDIR /app
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./main.go ./main.go
COPY ./bin ./bin
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./apps ./apps
COPY ./Makefile ./Makefile
RUN apk add --update make
RUN make build.bin

FROM golang:1.18.10-alpine3.17 AS RUN
ENV SUBCOMMAND=${SUBCOMMAND:-"start"}
WORKDIR /app
COPY --from=BUILD_GOLANG /app/bin /app/bin
CMD exec /app/bin/workshop ${SUBCOMMAND}
