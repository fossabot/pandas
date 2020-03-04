FROM ubuntu:18.04

LABEL vendor=cloudtone
LABEL product=pandas

RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list

#ENV DISTRO_NAME=zookeeper-3.4.14

ADD docker/install/postgressql.sh /install/
RUN chmod 755 install/postgressql.sh
RUN ./install/postgressql.sh
ADD conf/postgressqld.conf /etc/
ADD enterpoint/postgressql.sh /enterpoint/
RUN chmod 755 /enterpoint/postgressql.sh

ADD docker/install/rabbitmq.sh /install/
RUN chmod 755 install/rabbitmq.sh
RUN ./install/rabbitmq.sh
ADD conf/rabbitmq.conf /etc/rabbitmq/
ADD enterpoint/rabbitmq.sh /enterpoint/
RUN chmod 755 /enterpoint/rabbitmq.sh

ADD docker/install/etcd.sh /install/
RUN chmod 755 install/etcd.sh
RUN ./install/etcd.sh
ADD enterpoint/etcd.sh /enterpoint/
RUN chmod 755 /enterpoint/etcd.sh

# postgressql volume
VOLUME /var/lib/postgressqldb

# rabbit volume
VOLUME /var/lib/rabbitmq

#COPY docker-entrypoint.sh /usr/local/bin/
#ENTRYPOINT ["docker-entrypoint.sh"]

# postgressql port
EXPOSE 27017

# rabbitmq port
EXPOSE 4369 5671 5672 25672

# etcd port
EXPOSE 2181 2888 3888