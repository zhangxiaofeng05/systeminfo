ARG GO_VERSION=1.18
FROM golang:${GO_VERSION}-alpine AS builder
ARG PROJECT_NAME=systeminfo
ENV PROJECT_NAME=${PROJECT_NAME}
COPY . /usr/src/${PROJECT_NAME}
WORKDIR /usr/src/${PROJECT_NAME}
ENV GOPROXY=https://proxy.golang.com.cn,direct
RUN go build -o ./bin/server -v ./cmd/server

FROM golang:${GO_VERSION}-alpine
ARG PROJECT_NAME=systeminfo
ENV PROJECT_NAME=${PROJECT_NAME}
WORKDIR /usr/src/${PROJECT_NAME}
COPY --from=builder /usr/src/${PROJECT_NAME}/bin/server .
EXPOSE 8080
ENTRYPOINT /usr/src/${PROJECT_NAME}/server
