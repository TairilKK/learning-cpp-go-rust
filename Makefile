CXX = g++
CXXFLAGS = -Wall -pedantic
RUSTC = rustc
RUSTFLAGS = -O
GO = go
GOFLAGS = build

BIN_DIR = ./bin
SRC_DIR = .

CPP_PROBLEMS := $(wildcard problem_*/main.cpp)
RUST_PROBLEMS := $(wildcard problem_*/main.rs)
GO_PROBLEMS := $(wildcard problem_*/main.go)

CPP_TARGETS := $(patsubst %/main.cpp,%,$(CPP_PROBLEMS))
RUST_TARGETS := $(patsubst %/main.rs,%,$(RUST_PROBLEMS))
GO_TARGETS := $(patsubst %/main.go,%,$(GO_PROBLEMS))

all: $(CPP_TARGETS) $(RUST_TARGETS) $(GO_TARGETS)

LAST_CPP := $(lastword $(sort $(CPP_PROBLEMS)))
LAST_RUST := $(lastword $(sort $(RUST_PROBLEMS)))
LAST_GO := $(lastword $(sort $(GO_PROBLEMS)))

LAST_CPP_TARGET := $(patsubst %/main.cpp,%,$(LAST_CPP))
LAST_RUST_TARGET := $(patsubst %/main.rs,%,$(LAST_RUST))
LAST_GO_TARGET := $(patsubst %/main.go,%,$(LAST_GO))

last_all: $(LAST_CPP_TARGET) $(LAST_RUST_TARGET) $(LAST_GO_TARGET)
last_cpp: $(LAST_CPP_TARGET)
last_rust: $(LAST_RUST_TARGET)
last_go: $(LAST_GO_TARGET)

$(CPP_TARGETS): %: %/main.cpp
	@mkdir -p $(BIN_DIR)
	$(CXX) $< -o $(BIN_DIR)/$(notdir $@).exe $(CXXFLAGS)

$(RUST_TARGETS): %: %/main.rs
	@mkdir -p $(BIN_DIR)
	$(RUSTC) $(RUSTFLAGS) $< -o $(BIN_DIR)/$(notdir $@).exe

$(GO_TARGETS): %: %/main.go
	@mkdir -p $(BIN_DIR)
	cd $(dir $<) && $(GO) $(GOFLAGS) -o ../$(BIN_DIR)/$(notdir $@).exe

clean:
	rm -f $(BIN_DIR)/*.exe

.PHONY: all clean $(CPP_TARGETS) $(RUST_TARGETS) $(GO_TARGETS)
