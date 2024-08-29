CC       = go build
CFLAGS   = -trimpath
LDFLAGS  = all=-w -s
GCFLAGS  = all=
ASMFLAGS = all=

all: client server

.PHONY: client
client:
	GO111MODULE=$(GOMOD) \
	GOARCH=$(GO_ARCH_AMD) \
	GOOS=$(GO_OS_LINUX) \
	CGO_ENABLED=0 \
	$(CC) $(CFLAGS) -ldflags="$(LDFLAGS)" -gcflags="$(GCFLAGS)" -asmflags="$(ASMFLAGS)" \
	-o ./client ./cmd/client

.PHONY: server
server:
	GO111MODULE=$(GOMOD) \
	GOARCH=$(GO_ARCH_AMD) \
	GOOS=$(GO_OS_LINUX) \
	CGO_ENABLED=1 \
	$(CC) $(CFLAGS) -ldflags="$(LDFLAGS)" -gcflags="$(GCFLAGS)" -asmflags="$(ASMFLAGS)" \
	-o ./server ./cmd/server
