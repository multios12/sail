#!/bin/sh

set -eu

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cd "$SCRIPT_DIR"

if ! command -v yarn >/dev/null 2>&1; then
  echo "yarn is required for the front-end build" >&2
  exit 1
fi

if ! command -v go >/dev/null 2>&1; then
  echo "go is required for the back-end build" >&2
  exit 1
fi

(
  cd front
  yarn build
)

mkdir -p api/cmd/sail/static
cp ./front/dist/* ./api/cmd/sail/static/ -R

(
  cd api
  go build ./...
)
