# Don't allow an implicit 'all' rule. This is not a user-facing file.
ifeq ($(MAKECMDGOALS),)
    $(error This Makefile requires an explicit rule to be specified)
endif

ifeq ($(DBG_MAKEFILE),1)
    $(warning ***** starting Makefile.generated_files for goal(s) "$(MAKECMDGOALS)")
    $(warning ***** $(shell date))
endif

# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /bin/bash
ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"
# We don't need make's built-in rules.
MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

# Constants used throughout.
.EXPORT_ALL_VARIABLES:
OUT_DIR ?= _output
BIN_DIR := $(OUT_DIR)/bin

.PHONY: build update clean

build: bazel-build

simple-build:
	bazel build --watchfs -- //tools/... -//vendor/...
ifeq ($(WHAT),)
bazel-build:
	bazel build --watchfs  -- //...
else 
bazel-build: 
	bazel build --watchfs -- //$(WHAT)/...
endif

build-keep-going:
	bazel build --watchfs -k -- //...
clean:
	bazel clean --expunge
update:
	kazel
test:
	@if [ "$(WHAT)" !=  "" ]; \
         then \
	 cd $(WHAT) && make ; \
	 else \
	 echo "Please input the WHAT" ;\
	 fi

bazel-test:
	@if [ "$(WHAT)" !=  "" ]; \
         then \
	 bazel test --watchfs -- //$(WHAT)/... ; \
	 else \
	 echo "Please input the WHAT" ;\
	 fi
check:
	@./hack/check.sh
init:
	go get github.com/gogo/protobuf/proto
	go get github.com/gogo/protobuf/gogoproto
	go get github.com/gogo/protobuf/jsonpb
	go get google.golang.org/grpc
	go get github.com/golang/oauth2
	go get cloud.google.com/go
