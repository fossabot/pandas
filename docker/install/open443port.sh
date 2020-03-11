#!/bin/bash

apt-get update
apt-get install sudo
apt-get -y install iptables

sudo iptables -I INPUT -p tcp --dport 443 -j ACCEPT

sudo iptables-save

