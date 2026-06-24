# 多阶段构建：编译静态二进制 → 跑在极小镜像里
FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
# 纯标准库，无需 go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/placeholder .

FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/placeholder /app/placeholder
EXPOSE 8080
ENTRYPOINT ["/app/placeholder"]
