#!/bin/bash
set -e

source $(realpath $(dirname $0)/build.env)

echo "${BUILD_OUT}:"
echo -ne "  Buidling..."
CGO_ENABLED=0 GOOS=${BUILD_OS} GOARCH=${BUILD_ARCH} go build -o ${BUILD_OUT} -ldflags "\
    -X '${GO_BUILDMETA}.GitTag=${GIT_TAG}' \
    -X '${GO_BUILDMETA}.GitCommit=${GIT_COMMIT}' \
    -X '${GO_BUILDMETA}.GitBranch=${GIT_BRANCH}' \
    -X '${GO_BUILDMETA}.BuildDate=${BUILD_DATE}' \
    -X '${GO_BUILDMETA}.Platform=${BUILD_PLATFORM}' \
    -extldflags '-static'"
echo -ne " ✔\n"

echo -ne "  Compressing..."
cd dist && tar -cvzf ${BUILD_BIN}.tar.gz ${BUILD_BIN} > /dev/null 2>&1 && cd ..
echo -ne " ✔\n"

echo -ne "  Calculating sha256..."
echo $(gsha256sum "${BUILD_OUT}.tar.gz") | cut -d " " -f1 > "${BUILD_OUT}tar.gz.sha256"
echo -ne " ✔\n"
