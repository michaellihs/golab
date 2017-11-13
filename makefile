targets := $(wildcard *.go)

compile: $(targets)
	go install github.com/michaellihs/golab

gendoc: $(targets)
	golab gendoc -p doc
