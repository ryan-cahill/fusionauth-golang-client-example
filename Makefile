SHELL := /bin/bash
.PHONY: all .check-env-vars build tag push

build:
	docker build . -t fa-golang-example

run:
	docker run -p $(PUBLIC_PORT):$(PUBLIC_PORT) --env-file=.env fa-golang-example
