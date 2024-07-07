package main

import (
	"github.com/johnfercher/chaos/arch/archservices"
	"github.com/johnfercher/chaos/struct/structservices"
	"log"
)

func main() {
	file := structservices.NewFile()
	classifier := structservices.NewFileClassifier()
	interfaceInterpreter := structservices.NewInterfaceInterpreter()
	functionInterpreter := structservices.NewFunctionInterpreter()
	discover := structservices.NewDiscover(file, classifier, interfaceInterpreter, functionInterpreter)
	importsDiscover := archservices.NewImportsDiscover()

	path := "docs/examples/go-hexagonal"
	files, err := discover.Project(path)
	if err != nil {
		log.Fatal(err)
	}

	_ = importsDiscover.Find(files)
}
