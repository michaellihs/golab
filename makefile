targets := $(wildcard *.go)

compile: $(targets)
	go install github.com/michaellihs/golab

gendoc: compile
	golab gendoc -p doc

test: compile
	go test ./tests -v
