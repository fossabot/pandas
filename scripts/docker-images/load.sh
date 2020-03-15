#!/usr/bin/env bash

docker load < elcolio_etcd.tar
docker load < bitnami_rabbitmq.tar
docker load < redis.tar
# docker load < janeczku_go-dnsmasq.tar
