#!/bin/bash

#apk update
#apk add ca-certificates
#update-ca-certificates
#apk --no-cache add openssl wget

apt-get update
apt-get -y install wget build-essential

wget https://www.sqlite.org/2019/sqlite-autoconf-3280000.tar.gz

tar -xzvf sqlite-autoconf-3280000.tar.gz
cd sqlite-autoconf-3280000
./configure --prefix=/usr/local
make && make install
ln -s /usr/local/bin/sqlite3 /usr/bin/sqlite3
export LD_LIBRARY_PATH="/usr/local/bin"
