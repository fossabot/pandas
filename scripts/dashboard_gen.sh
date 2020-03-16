#!/bin/bash
set -e -o pipefail

: "${WORKDIR:=./dashboard/}"
: "${NOCOMPRESS:=false}"

YARN_INSTALL="pushd ${WORKDIR} && \
                yarn && \
                popd"

YARN_BUILD_PROD="pushd ${WORKDIR} && \
                yarn build:prod && \
                popd"

YARN_BUILD_SIT="pushd ${WORKDIR} && \
                yarn build:stage && \
                popd"

GO_BINDATA="pushd ${WORKDIR} && \
                go-bindata-assetfs -pkg dashboard -nocompress=${NOCOMPRESS} dist/... && \
                popd"

bash -c "${YARN_INSTALL}"
bash -c "${YARN_BUILD_PROD}"
bash -c "${GO_BINDATA}"
