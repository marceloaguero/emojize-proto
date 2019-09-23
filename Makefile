# Copyright 2019 VMware, Inc.
# SPDX-License-Identifier: BSD-3-Clause

SHELL := /bin/bash

.PHONY: help
help:
	@echo "Usage: make <TARGET>"
	@echo ""
	@echo "Available targets are:"
	@echo ""
	@echo "    generate-go              Generate Go files from proto"
	@echo "    generate-js              Generate JavaScript files from proto"
	@echo ""

.PHONY: generate-go
generate-go:
	protoc -I ./ emoji.proto \
		--go_out=plugins=grpc:.

.PHONY: generate-js
generate-js:
	protoc -I ./ emoji.proto \
		--js_out=import_style=commonjs:. \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:.
