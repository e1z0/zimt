#!/bin/bash
set -e

platform=$1

if [[ -z $platform ]]; then
    echo "USAGE: PLATFORM" > /dev/stderr
    exit 1
fi

platform_parts=(${platform//\// })
os=${platform_parts[0]}
arch=${platform_parts[1]}
out="dist/zimt_${os}_${arch}"

module=$(awk '/module/{print $2}' go.mod)
buildmeta="${module}/pkg/buildmeta"

echo "Buidling ${out}..."
GOOS=${os} GOARCH=${arch} go build -o ${out} -ldflags "\
    -X '${buildmeta}.GitTag=$(git describe --abbrev=0)' \
    -X '${buildmeta}.GitCommit=$(git rev-parse --verify --short HEAD)' \
    -X '${buildmeta}.GitBranch=$(git symbolic-ref --short -q HEAD)' \
    -X '${buildmeta}.BuildDate=$(date -u +"%Y-%m-%dT%H:%M:%SZ")' \
    -X '${buildmeta}.Platform=${platform}'"
