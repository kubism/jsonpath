GO         ?= go
LINTER     ?= golangci-lint
BUILDFLAGS += -installsuffix cgo --tags release

BUILD_PATH ?= $(shell pwd)

CMD = $(BUILD_PATH)/jp
CMD_SRC = cmd/*.go

all: lint clean build

.PHONY: build lint clean

clean:
	rm -f $(CMD)

lint:
	$(LINTER) run -v --no-config --deadline=5m

$(CMD):
	CGO_ENABLED=0 GOOS=linux $(GO) build -o $(CMD) -a $(BUILDFLAGS) $(LDFLAGS) $(CMD_SRC)

build: $(CMD) 