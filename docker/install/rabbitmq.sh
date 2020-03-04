#!/bin/bash

# 安装erlangy语言支持
apt-get -y install erlang-nox

# 添加公钥
#wget -O- https://www.rabbitmq.com/rabbitmq-release-signing-key.asc | sudo apt-key add -

# 更新软件包
apt-get -y update

# 安装RabbitMq
apt-get -y install rabbitmq-server

#start
service rabbitmq-server start

