# 基础镜像
FROM golang:1.18-alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 工作目录路径
WORKDIR /project/go-docker/

# 从项目复制go.mod 和 go.sum 文件到工作目录
COPY go.* ./

# 从 go modules 官方镜像获取依赖
RUN go mod download

# 将我们项目中的所有内容复制到工作目录中
COPY . .

# 在工作目录中执行 go build 命令生成二进制文件
RUN go build -o /project/go-docker/build/myapp .

# 告诉 docker 我们的代码将暴露端口 8080
EXPOSE 8080

# 当我们运行这个镜像容器时，它将从我们的构建可执行文件开始执行
ENTRYPOINT [ "/project/go-docker/build/myapp" ]
