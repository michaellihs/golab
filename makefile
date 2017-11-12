compile: cmd/*.go
	go install github.com/michaellihs/golab

renderdoc: cmd/*.go
	golab gendoc -p doc
