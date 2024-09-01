DIR = ./build
GO_ARCH_AMD = amd64
GO_ARCH_ARM = arm64
GO_OS_LINUX = linux

EXECUTABLE_PREFIX = ftsync

BIN_LINUX_SERVER_AMD64 = $(DIR)/$(EXECUTABLE_PREFIX)-server-$(GO_ARCH_AMD)
BIN_LINUX_SERVER_ARM64 = $(DIR)/$(EXECUTABLE_PREFIX)-server-$(GO_ARCH_ARM)

BIN_LINUX_CLIENT_AMD64 = $(DIR)/$(EXECUTABLE_PREFIX)-client-$(GO_ARCH_AMD)
BIN_LINUX_CLIENT_ARM64 = $(DIR)/$(EXECUTABLE_PREFIX)-client-$(GO_ARCH_ARM)

LDFLAGS = -extldflags=-static

.PHONY:
all: client server

.PHONY: client
client:
	GOARCH=$(GO_ARCH_AMD) GOOS=$(GO_OS_LINUX) CGO_ENABLED=0 go build -o $(BIN_LINUX_CLIENT_AMD64) ./cmd/client
	GOARCH=$(GO_ARCH_ARM) GOOS=$(GO_OS_LINUX) CGO_ENABLED=0 go build -o $(BIN_LINUX_CLIENT_ARM64) ./cmd/client

.PHONY: server
server:
	GOARCH=$(GO_ARCH_AMD) GOOS=$(GO_OS_LINUX) CGO_ENABLED=$(CGO_ENABLED) \
	go build -ldflags="$(LDFLAGS)" -tags osusergo,netgo,sqlite_omit_load_extension -o $(BIN_LINUX_SERVER_AMD64) ./cmd/server

	GOARCH=$(GO_ARCH_ARM) GOOS=$(GO_OS_LINUX) GOARM=7 CGO_ENABLED=$(CGO_ENABLED) \
	CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ \
	go build -ldflags="$(LDFLAGS)" -tags osusergo,netgo,sqlite_omit_load_extension -o $(BIN_LINUX_SERVER_ARM64) ./cmd/server

clean:
	rm $(DIR)/*
