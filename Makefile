GO ?= go

BUILD_DIR ?= build
BIN_DIR = $(BUILD_DIR)/bin

MOD_BASE_PKG = github.com/RuchDB/chaos

CMD_DIR = cmd
CMDS = $(patsubst $(CMD_DIR)/%/,%,$(dir $(wildcard $(CMD_DIR)/*/.)))
CMD_BUILDS = $(patsubst %,%.build,$(CMDS))


.PHONY: all
all: build test


.PHONY: pre-built
pre-built:
	@mkdir -p $(BIN_DIR)

.PHONY: build
build: $(CMD_BUILDS)

.PHONY: $(CMD_BUILDS)
$(CMD_BUILDS): %.build: pre-built
	$(GO) build -o $(BIN_DIR)/$* $(MOD_BASE_PKG)/$(CMD_DIR)/$*


.PHONY: test
test:
	$(GO) test ./...


.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
