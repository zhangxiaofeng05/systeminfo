FROM golang:1.22-alpine AS builder
COPY . /app
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
RUN CGO_ENABLED=0 go build -ldflags '-w -s'

# FROM alpine:latest # if need cgo, use this
# FROM gcr.io/distroless/static
FROM scratch
COPY --from=builder /app/systeminfo /
EXPOSE 8080
ENTRYPOINT [ "/systeminfo" ]
