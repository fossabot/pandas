#!/bin/bash
protoc -I. -I../../ --gobson_out=plugins:. *.proto
cp ./dmms.proto ../../pms/grpc_pms_v1/github.com/cloustone/pandas/dmms/grpc_dmms_v1
