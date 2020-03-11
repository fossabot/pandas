FROM ubuntu:18.04

RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list

LABEL vendor=cloustone
LABEL product=pandas

ENV PATH=$PATH:/usr/local/protobuf/bin/
ENV PKG_CONFIG_PATH=/usr/local/protobuf/lib/pkgconfig/

ADD certs/* /etc/ssl/pandas
ADD bin/* /
RUN chmod 755 pandas-apimachinery

ADD docker/install/protoc.sh /install/
RUN chmod 755 install/protoc.sh
RUN ./install/protoc.sh
# RUN chmod 755 /dockerize

USER root

CMD ["/pandas-apimachinery"]
