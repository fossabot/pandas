FROM ubuntu:18.04 

LABEL vendor=cloustone
LABEL product=pandas

#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ENV PATH=$PATH:/usr/local/protobuf/bin/
ENV PKG_CONFIG_PATH=/usr/local/protobuf/lib/pkgconfig/

ADD certs/* /etc/ssl/pandas
ADD bin/* /
RUN chmod 755 pandas-headmast
#RUN chmod 755 /dockerize
ADD docker/install/protoc.sh /install/
RUN chmod 755 install/protoc.sh
RUN ./install/protoc.sh

USER root

CMD ["/pandas-headmast"]
