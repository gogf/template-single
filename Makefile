ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "${MODULE_NAME}"
DOCKER_NAME = "${MODULE_NAME}"

include ./hack/hack.mk