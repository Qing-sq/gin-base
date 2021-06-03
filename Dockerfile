FROM golang:alpine AS  builder

# 为我们的的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器
COPY . .

# 将我们代码编译成二进制可执行文件app
RUN go build -o app .

#==

# 移动到用于存放生成的二进制文件 /dist目录
WORKDIR /dist

# 二进制文件复制
RUN cp /build/app .

EXPOSE 8888