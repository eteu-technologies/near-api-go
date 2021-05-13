#!/usr/bin/env bash
set -euo pipefail

docker-compose exec nearup tail -f /root/.nearup/logs/localnet/node0.log
