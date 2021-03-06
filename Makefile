GO         ?= go
LINTER     ?= golangci-lint
BUILDFLAGS += -installsuffix cgo --tags release

BUILD_PATH ?= $(shell pwd)

CMD = $(BUILD_PATH)/jp
CMD_SRC = cmd/*.go

all: lint clean build

.PHONY: build lint clean

clean:
	rm -f $(CMD).linux
	rm -f $(CMD).darwin

lint:
	$(LINTER) run -v --no-config --deadline=5m

$(CMD):
	CGO_ENABLED=0 GOOS=linux $(GO) build -o $(CMD).linux -a $(BUILDFLAGS) $(LDFLAGS) $(CMD_SRC)
	CGO_ENABLED=0 GOOS=darwin $(GO) build -o $(CMD).darwin -a $(BUILDFLAGS) $(LDFLAGS) $(CMD_SRC)

build: $(CMD)