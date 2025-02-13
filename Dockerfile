# 第一阶段：构建阶段
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /app

# 将源代码复制到工作目录
COPY . .

# 编译 Go 应用
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# 第二阶段：运行阶段
FROM alpine:3.8

# 设置工作目录
WORKDIR /app

# 从构建阶段复制二进制文件和资源
COPY --from=builder /app/main ./main
COPY --from=builder /app/manifest/config ./manifest/config
COPY --from=builder /app/resource ./resource

# 确保二进制文件具有执行权限
RUN chmod +x ./main

# 设置默认命令
CMD ["./main"]