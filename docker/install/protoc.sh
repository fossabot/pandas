#!/bin/bash

apt-get update
apt-get -y install wget build-essential

wget http://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protobuf-all-3.11.4.tar.gz
tar xzvf protobuf-all-3.11.4.tar.gz
cd protobuf-3.11.4

./configure --prefix=/usr/local/protobuf

make && make install

export PATH=$PATH:/usr/local/protobuf/bin/
export PKG_CONFIG_PATH=/usr/local/protobuf/lib/pkgconfig/

protoc --version

