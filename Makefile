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

SORTED_CPP := $(sort $(CPP_PROBLEMS))
COUNT_CPP := $(words $(SORTED_CPP))
SECOND_TO_LAST_INDEX := $(shell echo $$(($(COUNT_CPP) - 1)))
SECOND_TO_LAST_CPP := $(word $(SECOND_TO_LAST_INDEX), $(SORTED_CPP))

# LAST_CPP_TARGET := $(patsubst %/main.cpp,%,$(LAST_CPP))
LAST_CPP_TARGET := $(patsubst %/main.cpp,%,$(SECOND_TO_LAST_CPP))
LAST_RUST_TARGET := $(patsubst %/main.rs,%,$(LAST_RUST))
LAST_GO_TARGET := $(patsubst %/main.go,%,$(LAST_GO))

last_cpp: $(LAST_CPP_TARGET)
last_rust: $(LAST_RUST_TARGET)
last_go: $(LAST_GO_TARGET)
last_all: $(LAST_CPP_TARGET) $(LAST_RUST_TARGET) $(LAST_GO_TARGET)

$(CPP_TARGETS): %: %/main.cpp
	@mkdir -p $(BIN_DIR)
	$(CXX) $< -o $(BIN_DIR)/$(notdir $@).exe $(CXXFLAGS)

$(RUST_TARGETS): %: %/main.rs
	@mkdir -p $(BIN_DIR)
	$(RUSTC) $(RUSTFLAGS) $< -o $(BIN_DIR)/$(notdir $@).exe

$(GO_TARGETS): %: %/main.go
	@mkdir -p $(BIN_DIR)
	cd $(dir $<) && $(GO) $(GOFLAGS) -o ../$(BIN_DIR)/$(notdir $@).exe

run_last_cpp: $(LAST_CPP_TARGET)
	@echo "Running latest C++ program: $(BIN_DIR)/$(notdir $(LAST_CPP_TARGET)).exe"
	@./$(BIN_DIR)/$(notdir $(LAST_CPP_TARGET)).exe

run_last_rust: $(LAST_RUST_TARGET)
	@echo "Running latest Rust program: $(BIN_DIR)/$(notdir $(LAST_RUST_TARGET)).exe"
	@./$(BIN_DIR)/$(notdir $(LAST_RUST_TARGET)).exe

run_last_go: $(LAST_GO_TARGET)
	@echo "Running latest Go program: $(BIN_DIR)/$(notdir $(LAST_GO_TARGET)).exe"
	@./$(BIN_DIR)/$(notdir $(LAST_GO_TARGET)).exe

clean:
	rm -f $(BIN_DIR)/*.exe

.PHONY: all clean $(CPP_TARGETS) $(RUST_TARGETS) $(GO_TARGETS)
