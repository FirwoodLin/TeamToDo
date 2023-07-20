# docker中的包名
FROM golang:default

# 当前打包的工作路径
WORKDIR /go/src

# 赋值工作路径下的代码到容器中的工作目录
COPY . .

# 打开模块开关
RUN go env -w GO111MODULE="on"
# 设置代理
RUN go env -w GOPROXY="https://goproxy.cn,direct"

# go build时会去go.mod中去看是否有缺失模块，如果有就自动获取
# 在当前目录下对文件进行编译
RUN go build -o app

# 暴露指定端口
EXPOSE 8080

# 开始运行程序
CMD ["./app"]


