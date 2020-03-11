FROM ubuntu:18.04

LABEL vendor=cloustone
LABEL product=pandas

ENV PATH=$PATH:/usr/local/protobuf/bin/
ENV PKG_CONFIG_PATH=/usr/local/protobuf/lib/pkgconfig/

ADD certs/* /etc/ssl/pandas
ADD bin/* /
RUN chmod 755 pandas-lbs
# RUN chmod 755 /dockerize

ADD docker/install/open443port.sh /install/
RUN chmod 755 install/open443port.sh
RUN ./install/open443port.sh

ADD docker/install/protoc.sh /install/
RUN chmod 755 install/protoc.sh
RUN ./install/protoc.sh

USER root

CMD ["/pandas-lbs"]
