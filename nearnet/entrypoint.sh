#!/usr/bin/env bash
set -euo pipefail

datadir=/root/.near
pidfile="${HOME}/.nearup/node.pid"

if [ -f "${pidfile}" ]; then
	rm "${pidfile}"
fi

nearup "${@}"
sleep 1

pid="$(sed 's/^\(\w\+\)|.*/\1/g' "${pidfile}" | head -1)"
while (kill -0 "${pid}" &>/dev/null); do
	sleep 2
done
