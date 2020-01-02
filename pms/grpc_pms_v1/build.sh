#!/bin/bash
#protoc -I. -I../../dmms/grpc_dmms_v1/ -I../../pkg/proto --gobson_out=plugins:. *.proto
protoc -I. -I../../pkg/proto --gobson_out=plugins:. *.proto
