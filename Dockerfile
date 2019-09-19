# 指定哪种镜像作为新镜像的基础镜像
FROM golang:1.12.6
#FROM alpine:latest
# 指明该镜像的作者和其电子邮件
MAINTAINER xiongdun "1274328268@qq.com"
# Create the directory where the application will reside
#RUN mkdir /app
#
## Copy the application files (needed for production)
#ADD whacos/whacos /app/whacos
#
## Set the working directory to the app directory
#WORKDIR /app
#WORKDIR /usr/local/go/whacos
#WORKDIR $GOPATH/src/whacos
#ADD . $GOPATH/src/whacos
#COPY . .
#RUN cd /app/whacos && go build .
 # 编译一个静态的go应用（在二进制构建中包含C语言依赖库）
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
#RUN go build main.go
# Expose the application on port 8080.
# This should be the same as in the app.conf file
#RUN go get github.com/xiongdun/whacos
#WORKDIR $GOPATH/src/github.com/xiongdun/whacos
WORKDIR $GOPATH/src/github.com/xiongdun/whacos
COPY . $GOPATH/src/github.com/xiongdun/whacos
RUN go build .

EXPOSE 8090
ENTRYPOINT ["./main"]

# Set the entry point of the container to the application executable
#ENTRYPOINT ["./main"]
#
#FROM debian:stretch
#
#ARG DEBIAN_FRONTEND=noninteractive
#ARG JAVA_VERSION=8r
#ARG JAVA_UPDATE=172
#ARG JAVA_BUILD=11
#ARG JAVA_PACKAGE=jdk
#ARG JAVA_HASH=a58eab1ec242421181065cdc37240b08
#
#ENV LANG C.UTF-8
#ENV JAVA_HOME=/opt/jdk
#ENV PATH=${PATH}:${JAVA_HOME}/bin
#
#RUN set -ex \
# && apt-get update \
# && apt-get -y install ca-certificates wget unzip \
# && wget -q --header "Cookie: oraclelicense=accept-securebackup-cookie" \
#         -O /tmp/java.tar.gz \
#         http://download.oracle.com/otn-pub/java/jdk/${JAVA_VERSION}u${JAVA_UPDATE}-b${JAVA_BUILD}/${JAVA_HASH}/${JAVA_PACKAGE}-${JAVA_VERSION}u${JAVA_UPDATE}-linux-x64.tar.gz \
# && CHECKSUM=$(wget -q -O - https://www.oracle.com/webfolder/s/digest/${JAVA_VERSION}u${JAVA_UPDATE}checksum.html | grep -E "${JAVA_PACKAGE}-${JAVA_VERSION}u${JAVA_UPDATE}-linux-x64\.tar\.gz" | grep -Eo '(sha256: )[^<]+' | cut -d: -f2 | xargs) \
# && echo "${CHECKSUM}  /tmp/java.tar.gz" > /tmp/java.tar.gz.sha256 \
# && sha256sum -c /tmp/java.tar.gz.sha256 \
# && mkdir ${JAVA_HOME} \
# && tar -xzf /tmp/java.tar.gz -C ${JAVA_HOME} --strip-components=1 \
# && wget -q --header "Cookie: oraclelicense=accept-securebackup-cookie;" \
#         -O /tmp/jce_policy.zip \
#         http://download.oracle.com/otn-pub/java/jce/${JAVA_VERSION}/jce_policy-${JAVA_VERSION}.zip \
# && unzip -jo -d ${JAVA_HOME}/jre/lib/security /tmp/jce_policy.zip \
# && rm -rf ${JAVA_HOME}/jar/lib/security/README.txt \
#       /var/lib/apt/lists/* \
#       /tmp/* \
#       /root/.wget-hsts