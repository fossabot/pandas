#!/usr/bin/env bash

docker save elcolio/etcd > elcolio_etcd.tar
docker save bitnami/rabbitmq > bitnami_rabbitmq.tar
docker save redis:alpine > redis.tar
# docker save janeczku/go-dnsmasq > janeczku_go-dnsmasq.tar

# gcr.io/google_containers/defaultbackend:1.4
# quay.io/kubernetes-ingress-controller/nginx-ingress-controller:0.16.2
# quay.io/external_storage/nfs-client-provisioner:v1.0
# prom/alertmanager:v0.14.0
# prom/node-exporter:v0.16.0
# prom/pushgateway
# jimmidyson/configmap-reload:v0.1
# k8s.gcr.io/fluentd-elasticsearch:v2.2.0
# docker.elastic.co/kibana/kibana-oss:6.2.4
