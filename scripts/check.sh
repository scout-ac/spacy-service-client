#!/usr/bin/env bash

set -euo pipefail

this_dir=$(readlink -qe $(cd -- $(dirname -- "${BASH_SOURCE[0]}") &>/dev/null && pwd))
wd=$(readlink -qe "${this_dir}"/../)

cd "${wd}"

{
	echo "Running revive..."
	revive -set_exit_status -formatter stylish ./...
	:
} && {
	echo "Running staticcheck..."
	staticcheck ./...
} && {
	echo "Running go vet..."
	go vet ./...
}
