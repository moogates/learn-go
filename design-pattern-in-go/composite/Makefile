SRC_FILES=$(wildcard *.go)

composite : $(SRC_FILES)
	go build -o $@ $(SRC_FILES) 

clean:
	rm -fv composite

.PHONY : clean
