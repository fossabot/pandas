FROM ubuntu:18.04

LABEL vendor=cloudtone
LABEL product=pandas

RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirrors.aliyun.com/g' /etc/apt/sources.list

ENV DISTRO_NAME=zookeeper-3.4.14

ADD install/mongo.sh /install/
RUN chmod 755 /install/mongo.sh
RUN install/mongo.sh
ADD conf/mongod.conf /etc/
ADD enterpoint/mongo.sh /enterpoint/
RUN chmod 755 /enterpoint/mongo.sh

ADD install/rabbitmq.sh /install/
RUN chmod 755 /install/rabbitmq.sh
RUN install/rabbitmq.sh
ADD conf/rabbitmq.conf /etc/rabbitmq/
ADD enterpoint/rabbitmq.sh /enterpoint/
RUN chmod 755 /enterpoint/rabbitmq.sh

ADD install/zookeeper.sh /install/
RUN chmod 755 /install/zookeeper.sh
RUN install/zookeeper.sh
ADD enterpoint/zookeeper.sh /enterpoint/
RUN chmod 755 /enterpoint/zookeeper.sh


# mongo volume
VOLUME /var/lib/mongodb

# rabbit volume
VOLUME /var/lib/rabbitmq

COPY docker-entrypoint.sh /usr/local/bin/
#ENTRYPOINT ["docker-entrypoint.sh"]

# mongo port
EXPOSE 27017

# rabbitmq port
EXPOSE 4369 5671 5672 25672

# zookeeper port
EXPOSE 2181 2888 3888
