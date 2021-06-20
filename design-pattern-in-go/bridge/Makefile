SRC_FILES=$(wildcard *.go)
main : $(SRC_FILES)
	go build -o $@ $(SRC_FILES)

clean:
	rm -fv main

.PHONY : clean
