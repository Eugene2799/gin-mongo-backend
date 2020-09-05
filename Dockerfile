#源镜像
FROM golang:latest
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
WORKDIR $GOPATH/src/gin-mongo-backend
COPY . $GOPATH/src/gin-mongo-backend
RUN go build .
#暴露端口
EXPOSE 5922
#最终运行docker的命令
ENTRYPOINT  ["./gin-mongo-backend"]
