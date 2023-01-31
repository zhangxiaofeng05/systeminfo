FROM golang:1.18-alpine AS builder
COPY . /app
WORKDIR /app
# ENV GOPROXY=https://proxy.golang.com.cn,direct
RUN go build

FROM golang:1.18-alpine
COPY --from=builder /app/systeminfo /
EXPOSE 8080
ENTRYPOINT /systeminfo
