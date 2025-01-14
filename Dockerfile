FROM registry.cn-zhangjiakou.aliyuncs.com/matrix-net/ubuntu:24.04
LABEL maintainer "matrix@126.com"

ENV TZ='Asia/Shanghai'

WORKDIR /matrix
COPY park ./
COPY swagger ./swagger
COPY conf/app-prd.conf    ./conf/app.conf
RUN chmod +x ./park
CMD ["/matrix/park"]