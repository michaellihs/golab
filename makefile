targets := $(wildcard *.go)

compile: $(targets)
	go install github.com/michaellihs/golab

gendoc: compile
	golab gendoc -p doc
	golab zsh-completion --path zsh/_golab

test: compile
	### run integration tests with Ginkgo
	cd cmd && ginkgo -v
	cd cmd/mapper && ginkgo -v
	### run acceptance tests against real instance
	go test ./tests -v
