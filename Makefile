BINARY := uvctl
MODULE := github.com/abhinand/uvctl

# Version info (overridden in CI)
VERSION ?= dev
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")

# Build flags
LDFLAGS := -s -w \
	-X '$(MODULE)/cmd.Version=$(VERSION)' \
	-X '$(MODULE)/cmd.Commit=$(COMMIT)'

# Platforms
PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64

# Output directory
DIST := dist

.PHONY: all build clean install test lint fmt vet $(PLATFORMS)

all: build

build:
	go build -ldflags="$(LDFLAGS)" -o $(BINARY) .

install: build
	mv $(BINARY) $(GOPATH)/bin/$(BINARY)

clean:
	rm -f $(BINARY)
	rm -rf $(DIST)

test:
	go test -v ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint: fmt vet

# Cross-compilation targets
$(PLATFORMS):
	@mkdir -p $(DIST)
	GOOS=$(word 1,$(subst /, ,$@)) GOARCH=$(word 2,$(subst /, ,$@)) \
		go build -ldflags="$(LDFLAGS)" -o $(DIST)/$(BINARY)-$(word 1,$(subst /, ,$@))-$(word 2,$(subst /, ,$@)) .

release: $(PLATFORMS)
