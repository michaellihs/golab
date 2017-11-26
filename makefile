targets := $(wildcard *.go)

compile: $(targets)
	go install github.com/michaellihs/golab

gendoc: compile
	golab gendoc -p doc
	golab zsh-completion --path zsh/_golab

test: compile
	go test ./tests -v
	# cd cmd && ginkgo -v
