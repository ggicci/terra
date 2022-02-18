default: build

SHELL=/usr/bin/env bash
export DOCKER_UID=$(shell id -u)
export DOCKER_GID=$(shell id -g)

build:
	go build -o bin/terra

start:
	IO4_TERRA_ENV=dev bin/terra

dev: dev/build-dev-image dev/start-dev-container

dev/build-dev-image: assert/on-host # build dev image
	docker buildx build -f ci/dev/Dockerfile -t dev-terra:latest .

dev/start-dev-container: assert/on-host # start a container named "dev-terra" as the dev environment
	mkdir -p ${HOME}/.io4/terra/dev/{home,pg-data}
	docker compose create
	docker compose start

assert/on-host:
	@if [[ "${IO4_TERRA_DEV_IN_CONTAINER}" = "1" ]]; then \
		>&2 echo "[WARN] \`make dev\` is only allowed on the host, you are in the container."; \
		exit 1; \
	fi

.PHONY: build dev
