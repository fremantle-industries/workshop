FROM node:19-alpine3.15 AS BUILD_NODEJS
WORKDIR /app
COPY ./pkg ./pkg
COPY ./Makefile ./Makefile
RUN apk add --update make
RUN make setup.deps.controlui
RUN make build.controlui

FROM golang:1.19.3-alpine3.16 AS BUILD_GOLANG
WORKDIR /app
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./main.go ./main.go
COPY ./bin ./bin
COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./workshop.example.yaml ./workshop.yaml
COPY ./Makefile ./Makefile
COPY --from=BUILD_NODEJS /app/pkg/control/ui/dist ./pkg/control/ui/dist
RUN apk add --update make
RUN make build.bin

FROM golang:1.19.3-alpine3.16 AS RUN
ENV SUBCOMMAND=${SUBCOMMAND:-"--version"}
ENV HTTP_PORT=${HTTP_PORT:-8080}
WORKDIR /app
COPY --from=BUILD_GOLANG /app/bin /app/bin
CMD exec /app/bin/workshop ${SUBCOMMAND}
