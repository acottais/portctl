# I usually keep a `VERSION` file in the root so that anyone
# can clearly check what's the VERSION of `master` or any
# branch at any time by checking the `VERSION` in that git
# revision.
#
# Another benefit is that we can pass this file to our Docker 
# build context and have the version set in the binary that ends
# up inside the Docker image too.
VERSION           :=  $(shell cat ./VERSION)
IMAGE_NAME        :=  acottais/portctl
BUILD_DATE        :=  $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

DEP_VERSION       :=  0.5.3
BUILD_DIR         :=  build
PORTAINER_VERSION :=  1.20.2
DOCKER_VERSION    :=  1.40

GO_LDFLAGS      := "-X github.com/acottais/portctl/cmd.Version=$(VERSION) \
-X github.com/acottais/portctl/cmd.BuildDate=$(BUILD_DATE) \
-X github.com/acottais/portctl/cmd.PortainerVersion=$(PORTAINER_VERSION) \
-X github.com/acottais/portctl/cmd.DockerVersion=$(DOCKER_VERSION)"

# As a call to `make` without any arguments leads to the execution
# of the first target found I really prefer to make sure that this
# first one is a non-destructive one that does the most simple 
# desired installation. 
#
# It's very common to people set it as `all` but it could be anything 
# like `a`.
all: install

# install dep if not present
/usr/local/bin/dep: 
	curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v$(DEP_VERSION)/dep-linux-amd64
	chmod +x /usr/local/bin/dep

# Ensure dependencies are present
dep: /usr/local/bin/dep
	dep ensure


# Build the project
build: dep
	go build -o $(BUILD_DIR)/portctl -ldflags=$(GO_LDFLAGS)

# Install just performs a normal `go install` which builds the source
# files from the package at `./` (I like to keep a `main.go` in the root
# that imports other subpackages). 
#
# As I always commit `vendor` to `git`, a `go install` will typically 
# always work - except if there's an OS limitation in the build flags 
# (e.g, a linux-only project).
install: dep
	go install -v


# Keeping `./main.go` with just a `cli` and `./lib/*.go` with actual 
# logic, `tests` usually reside under `./lib` (or some other subdirectories).
#
# By using the `./...` notation, all the non-vendor packages are going
# to be tested if they have test files.
test:
	go test ./... -v


# Just like `test`, formatting what matters. As `main.go` is in the root,
# `go fmt` the root package. Then just `cd` to what matters to you (`vendor`
# doesn't matter).
#
# By using the `./...` notation, all the non-vendor packages are going
# to be formatted (including test files).
fmt:
	go fmt ./... -v


# This target is only useful if you plan to also create a Docker image at
# the end. 
#
# I really like publishing a Docker image together with the GitHub release
# because Docker makes it very simple to someone run your binary without
# having to worry about the retrieval of the binary and execution of it
# - docker already provides the necessary boundaries.
image:
	docker build -t $(IMAGE_NAME):$(VERSION) .


# This is pretty much an optional thing that I tend always to include.
#
# Goreleaser is a tool that allows anyone to integrate a binary releasing
# process to their pipelines. 
#
# Here in this target With just a simple `make release` you can have a 
# `tag` created in GitHub with multiple builds if you wish. 
#
# See more at `gorelease` github repo.
release:
	git tag -a $(VERSION) -m "Release $(VERSION)" || true
	git push origin $(VERSION)

.PHONY: install test fmt release dep build