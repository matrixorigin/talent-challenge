ROOT_DIR = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/
LD_FLAGS = -ldflags "-w -s"

GOOS 		= linux
DIST_DIR 	= $(ROOT_DIR)dist/

.PHONY: dist_dir
dist_dir: ; $(info ======== prepare distribute dir:)
	mkdir -p $(DIST_DIR)
	@rm -rf $(DIST_DIR)

.PHONY: server
server: dist_dir; $(info ======== compiled server)
	env GO111MODULE=on GOOS=$(GOOS) go build -mod=vendor -a -installsuffix cgo -o $(DIST_DIR)server $(LD_FLAGS) $(ROOT_DIR)cmd/server/*.go

.PHONY: checker
checker: dist_dir; $(info ======== compiled checker)
	env GO111MODULE=on GOOS=$(GOOS) go build -mod=vendor -a -installsuffix cgo -o $(DIST_DIR)checker $(LD_FLAGS) $(ROOT_DIR)cmd/checker/*.go

.PHONY: docker
docker: ; $(info ======== compiled baas docker)
	docker build -t matrixbase/server -f Dockerfile .

.DEFAULT_GOAL := server
