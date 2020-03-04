#!/bin/bash

# Install erlangy language support
apt-get -y install erlang-nox

# Add public key
#wget -O- https://www.rabbitmq.com/rabbitmq-release-signing-key.asc | sudo apt-key add -

# Update package
apt-get -y update

# Install RabbitMq
apt-get -y install rabbitmq-server

# start
service rabbitmq-server start

