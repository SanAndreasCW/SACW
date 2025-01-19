# Makefile for Go Project
# Variables
GOARCH=386
CGO_ENABLED=1
BUILD_OUTPUT=gamemodes/LSGW.so
RUN_COMMAND=./omp-server
DEV_COMMAND=nodemon -e .so --exec ${RUN_COMMAND}

.PHONY: all build dev run clean

all: build

build:
	GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) go build -buildmode=c-shared -o $(BUILD_OUTPUT)

dev:
	$(DEV_COMMAND)

run:
	$(RUN_COMMAND)

clean:
	@echo "Cleaning build files..."
	rm -f $(BUILD_OUTPUT)