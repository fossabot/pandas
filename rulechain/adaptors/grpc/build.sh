#!/bin/bash
protoc -I. -I../../ --gobson_out=plugins:. *.proto
