#!/bin/bash
set -e -o pipefail

: ${WORKDIR:="./dashboard/"}
: ${NOCOMPRESS:=false}

YARN_BUILD_PROD="pushd ${WORKDIR} && \
                yarn build:prod && \
                popd"

GO_BINDATA="pushd ${WORKDIR} && \
                go-bindata-assetfs -pkg dashboard -nocompress=${NOCOMPRESS} dist/... && \
                popd"

#bash -c "${YARN_BUILD_PROD}"
bash -c "${GO_BINDATA}"
