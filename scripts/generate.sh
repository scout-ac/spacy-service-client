#!/bin/bash

set -euo pipefail

this_dir=$(readlink -qe $(cd -- $(dirname -- "${BASH_SOURCE[0]}") &>/dev/null && pwd))
wd=$(readlink -qe "${this_dir}"/../)

cd "${wd}"

err_protoc_gen_go=0
err_protoc_gen_go_rpgc=0

command -v protoc-gen-go      > /dev/null || err_protoc_gen_go=1
command -v protoc-gen-go_grpc > /dev/null || err_protoc_gen_go_grpc=1

if [[ err_protoc_gen_go -ne 0 ]]; then
	echo "You need to install protoc-gen-go:"
	echo "    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
fi

if [[ err_protoc_gen_go -ne 0 ]]; then
	echo "You need to install protoc-gen-go-grpc:"
	echo "    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
fi

src="lib/spacy-service/proto"
dst="spacysvc/generated"

protoc \
	--proto_path="${src}" \
    --go_out="${dst}" \
    --go-grpc_out="${dst}" \
    --go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	"${src}"/spacy_service.proto
