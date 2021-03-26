#!/bin/sh

set -ex

# make check
dockerd >&2 2>/dev/null &
while ! docker info >/dev/null; do sleep 2; done

for state in ./target_state/*.yaml; do
	go run ensure_images.go -file $state
	compose-ref -f $state up
	docker ps -a
done
