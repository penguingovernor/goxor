GO_FILES := $(shell find . -name '*.go')
VERSION := 2.0.1

all: clean build

build: $(GO_FILES)
	go build -ldflags="-X github.com/penguingovernor/goxor/internal/constants.Version=${VERSION}"

clean:
	go clean