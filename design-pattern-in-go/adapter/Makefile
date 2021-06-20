SRC_FILES=$(wildcard *.go)
adapter : $(SRC_FILES)
	go build -o $@ $(SRC_FILES)

clean:
	rm -fv adapter

.PHONY : clean
