# Kompilatory i flagi
CXX = g++
CXXFLAGS = -Wall -pedantic
RUSTC = rustc
RUSTFLAGS = -O
GO = go
GOFLAGS = build

# Katalogi
BIN_DIR = ./bin
SRC_DIR = .

# Znajdź wszystkie problemy w różnych językach
CPP_PROBLEMS := $(wildcard problem_*/main.cpp)
RUST_PROBLEMS := $(wildcard problem_*/main.rs)
GO_PROBLEMS := $(wildcard problem_*/main.go)

# Generuj nazwy targetów (uproszczone)
CPP_TARGETS := $(patsubst %/main.cpp,%,$(CPP_PROBLEMS))
RUST_TARGETS := $(patsubst %/main.rs,%,$(RUST_PROBLEMS))
GO_TARGETS := $(patsubst %/main.go,%,$(GO_PROBLEMS))

# Domyślny target - kompiluje wszystko
all: $(CPP_TARGETS) $(RUST_TARGETS) $(GO_TARGETS)

# Reguła dla problemów C++
$(CPP_TARGETS): %: %/main.cpp
	@mkdir -p $(BIN_DIR)
	$(CXX) $< -o $(BIN_DIR)/$(notdir $@).exe $(CXXFLAGS)

# Reguła dla problemów Rust
$(RUST_TARGETS): %: %/main.rs
	@mkdir -p $(BIN_DIR)
	$(RUSTC) $(RUSTFLAGS) $< -o $(BIN_DIR)/$(notdir $@).exe

# Reguła dla problemów Go
$(GO_TARGETS): %: %/main.go
	@mkdir -p $(BIN_DIR)
	cd $(dir $<) && $(GO) $(GOFLAGS) -o ../$(BIN_DIR)/$(notdir $@).exe

# Czyszczenie
clean:
	rm -f $(BIN_DIR)/*.exe

.PHONY: all clean $(CPP_TARGETS) $(RUST_TARGETS) $(GO_TARGETS)
