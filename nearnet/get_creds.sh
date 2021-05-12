#!/usr/bin/env bash
set -euo pipefail

datadir="$(docker volume inspect nearnet_data --format '{{ .Mountpoint }}')"
sudo cat "${datadir}"/localnet/node0/validator_key.json
