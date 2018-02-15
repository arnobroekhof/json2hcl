#SHELL = bash
THIS_OS := $(shell uname)
CUR_DATE := $(shell DATE +%Y%m%d)
TARGET_NAME := json2hcl

ALL_TARGETS += linux-amd64 \
	linux-arm \
	linux-arm64 \
	windows-amd64 \
	darwin-amd64


clean:
	rm -rf build/*
	rm -rf vendor/*


deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

release:	clean	deps	build_binaries


build_binaries:
	$(foreach t, $(ALL_TARGETS), $(call go_build,$(t)))


define go_build
	$(shell export GOOS=$(shell cut -d'-' -f1 <<<$(1)); \
    export GOARCH=$(shell cut -d'-' -f2 <<<$(1)); \
	go build -o build/$(TARGET_NAME)-$(1)-$(CUR_DATE))
endef