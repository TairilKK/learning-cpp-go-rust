CXX = g++
CXXFLAGS = -Wall -pedantic
BIN_DIR = ./bin

all: problem_four

problem_one:
	$(CXX) ./Problem_1*/$@.cpp -o $(BIN_DIR)/$@.exe $(CXXFLAGS)

problem_two:
	$(CXX) ./Problem_2*/$@.cpp -o $(BIN_DIR)/$@.exe $(CXXFLAGS)

problem_three:
	$(CXX) ./Problem_3*/$@.cpp -o $(BIN_DIR)/$@.exe $(CXXFLAGS)

problem_four:
	$(CXX) ./Problem_4*/$@.cpp -o $(BIN_DIR)/$@.exe $(CXXFLAGS)

clean:
	rm -f $(BIN_DIR)/*.exe

.PHONY: all all_problems clean $(TARGETS)
