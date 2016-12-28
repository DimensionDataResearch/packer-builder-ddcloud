PLUGIN_NAME = ddcloud

VERSION = 0.1.0
VERSION_INFO_FILE = ./$(PLUGIN_NAME)/version-info.go

BIN_DIRECTORY   = _bin
EXECUTABLE_NAME = packer-builder-$(PLUGIN_NAME)
DIST_ZIP_PREFIX = $(EXECUTABLE_NAME).v$(VERSION)

REPO_BASE     = github.com/DimensionDataResearch
REPO_ROOT     = $(REPO_BASE)/packer-builder-ddcloud
PLUGIN_ROOT   = $(REPO_ROOT)/$(PLUGIN_NAME)
VENDOR_ROOT   = $(REPO_ROOT)/vendor

default: fmt build test

fmt:
	go fmt $(REPO_ROOT)/...

clean:
	rm -rf $(BIN_DIRECTORY) $(VERSION_INFO_FILE)
	go clean $(REPO_ROOT)/...

# Peform a development (current-platform-only) build.
# TODO: Copy plugin to ~/.packer.d/plugins.
dev: version fmt
	go build -o $(BIN_DIRECTORY)/$(EXECUTABLE_NAME)

# Perform a full (all-platforms) build.
build: version build-windows64 build-windows32 build-linux64 build-mac64

build-windows64: version
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIRECTORY)/windows-amd64/$(EXECUTABLE_NAME).exe

build-windows32: version
	GOOS=windows GOARCH=386 go build -o $(BIN_DIRECTORY)/windows-386/$(EXECUTABLE_NAME).exe

build-linux64: version
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIRECTORY)/linux-amd64/$(EXECUTABLE_NAME)

build-mac64: version
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIRECTORY)/darwin-amd64/$(EXECUTABLE_NAME)

# Build docker image
build-docker: build-linux64
	docker build -t ddresearch/packer-builder-$(PLUGIN_NAME) .
	docker tag ddresearch/packer-builder-$(PLUGIN_NAME) ddresearch/packer-builder-$(PLUGIN_NAME):v${VERSION}

# Build docker image
push-docker: build-docker
	docker push ddresearch/packer-builder-$(PLUGIN_NAME):latest
	docker push ddresearch/packer-builder-$(PLUGIN_NAME):v${VERSION}

# Produce archives for a GitHub release.
dist: build
	cd $(BIN_DIRECTORY)/windows-386 && \
		zip -9 ../$(DIST_ZIP_PREFIX).windows-386.zip $(EXECUTABLE_NAME).exe
	cd $(BIN_DIRECTORY)/windows-amd64 && \
		zip -9 ../$(DIST_ZIP_PREFIX).windows-amd64.zip $(EXECUTABLE_NAME).exe
	cd $(BIN_DIRECTORY)/linux-amd64 && \
		zip -9 ../$(DIST_ZIP_PREFIX).linux-amd64.zip $(EXECUTABLE_NAME)
	cd $(BIN_DIRECTORY)/darwin-amd64 && \
		zip -9 ../$(DIST_ZIP_PREFIX)-darwin-amd64.zip $(EXECUTABLE_NAME)

test: fmt # TODO: Add test targets

testall: 
	go test -v $(REPO_ROOT)/...

version: $(VERSION_INFO_FILE)

$(VERSION_INFO_FILE): Makefile
	@echo "Update version info: v$(VERSION)"
	@echo "package $(PLUGIN_NAME)\n\n// ProviderVersion is the current version of the $(PLUGIN_NAME) Builder plugin for Packer.\nconst ProviderVersion = \"v$(VERSION) (`git rev-parse HEAD`)\"" > $(VERSION_INFO_FILE)
