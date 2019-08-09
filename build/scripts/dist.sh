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
bin="zimt_${os}_${arch}"
out="dist/${bin}"

module=$(awk '/module/{print $2}' go.mod)
buildmeta="${module}/pkg/buildmeta"

echo "${out}:"
echo -ne "  Buidling..."
CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -o ${out} -ldflags "\
    -X '${buildmeta}.GitTag=$(git describe --abbrev=0)' \
    -X '${buildmeta}.GitCommit=$(git rev-parse --verify --short HEAD)' \
    -X '${buildmeta}.GitBranch=$(git symbolic-ref --short -q HEAD)' \
    -X '${buildmeta}.BuildDate=$(date -u +"%Y-%m-%dT%H:%M:%SZ")' \
    -X '${buildmeta}.Platform=${platform}' \
    -extldflags '-static'"
echo -ne " ✔\n"

echo -ne "  Compressing..."
cd dist && tar -cvzf ${bin}.tar.gz ${bin} > /dev/null 2>&1 && rm ${bin} && cd ..
echo -ne " ✔\n"

echo -ne "  Calculating sha256..."
echo $(gsha256sum "${out}.tar.gz") | cut -d " " -f1 > "${out}tar.gz.sha256"
echo -ne " ✔\n"
