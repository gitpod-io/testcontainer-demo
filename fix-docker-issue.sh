#!/bin/bash

tmpdir=$(mktemp -d)
cd "$tmpdir"
cat > Dockerfile <<EOF
FROM golang:1.21

WORKDIR /app
RUN git clone https://github.com/gitpod-io/gitpod && \
    cd gitpod && \
    git checkout 0a5dc029fc48b575c5a3fd1359f5eb751e88cef3
RUN cd gitpod/components/docker-up/runc-facade && go get && CGO_ENABLED=0 go build -o /app/runc-facade
EOF

docker build -t rfc .
docker run --rm -it -v $PWD:/dst rfc sh -c "cp /app/runc-facade /dst"
echo fixed runc-facade in "$PWD/runc-facade"
[ ! -f /usr/bin/runc-facade.org ] && sudo mv /usr/bin/runc-facade /usr/bin/runc-facade.org
sudo cp "$PWD/runc-facade" /usr/bin/runc-facade