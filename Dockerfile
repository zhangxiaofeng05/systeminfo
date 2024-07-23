FROM golang:1.22-alpine AS builder
# https://stackoverflow.com/questions/49118579/alpine-dockerfile-advantages-of-no-cache-vs-rm-var-cache-apk
RUN apk add --no-cache git
COPY . /app
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
RUN VERSION=$(git describe --tags --always --dirty) && \
    COMMIT=$(git rev-parse HEAD) && \
    CGO_ENABLED=0 go build -ldflags "-w -s -X main.version=$VERSION -X main.commit=$COMMIT"

# FROM alpine:latest # if need cgo, use this
# FROM gcr.io/distroless/static
FROM scratch
COPY --from=builder /app/systeminfo /
EXPOSE 8080
ENTRYPOINT [ "/systeminfo" ]
