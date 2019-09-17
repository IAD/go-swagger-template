#input
NAME ?= todo
SWAGGER_SPEC ?= swagger.yml
YAML_FILE := $(shell pwd)/examples/todo/${SWAGGER_SPEC}
CLIENT_PATH := $(shell pwd)/examples/todo/gen/client
CLIENT_PKG_SUFFIX ?= client

SERVER_PATH := $(shell pwd)/examples/todo/gen/server

#static
TEMPLATES_DIR := $(shell pwd)/templates

#dynamic
MODELS_PKG_SUFFIX ?= $(CLIENT_PKG_SUFFIX)models

gen-server:
	rm -rf $(SERVER_PATH)/*
	mkdir -p $(SERVER_PATH)
	docker run --rm -it -u `id -u $(USER)` \
		-e GOPATH=$(HOME)/go:/go \
		-v $(HOME):$(HOME) \
		-w $(shell pwd) quay.io/goswagger/swagger:v0.20.1 generate server \
		-f $(YAML_FILE) \
		-A $(NAME) \
		--template-dir ${TEMPLATES_DIR} \
		--exclude-main \
		-t $(SERVER_PATH)

gen-client:
	rm -rf $(CLIENT_PATH)/*
	mkdir -p $(CLIENT_PATH)
	docker run --rm -it -u `id -u $(USER)` \
		-e GOPATH=$(HOME)/go:/go \
		-v $(HOME):$(HOME) \
		-w $(shell pwd) quay.io/goswagger/swagger:v0.20.1 generate client \
		-f $(YAML_FILE) \
		-A $(NAME) \
		--template-dir ${TEMPLATES_DIR} \
		-c $(NAME)$(CLIENT_PKG_SUFFIX) \
		-m $(NAME)$(MODELS_PKG_SUFFIX) \
		-t $(CLIENT_PATH)
