FROM golang:1.18-alpine AS builder
COPY . /app
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
RUN CGO_ENABLED=0 go build -ldflags '-w -s'

FROM alpine:latest
COPY --from=builder /app/systeminfo /
EXPOSE 8080
ENTRYPOINT /systeminfo
