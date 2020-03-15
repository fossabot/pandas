#!/usr/bin/env bash

helm del --purge $(helm list | grep pandas- | awk '{print $1}')
kubectl delete svc $(kubectl get svc | grep pandas | awk '{print $1}')
kubectl delete deploy $(kubectl get deploy | grep pandas | awk '{print $1}')
