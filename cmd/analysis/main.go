package main

import (
	"github.com/johnfercher/chaos/struct/structservices"
	"log"
)

func main() {
	file := structservices.NewFile()
	classifier := structservices.NewFileClassifier()
	interfaceInterpreter := structservices.NewInterfaceInterpreter()
	functionInterpreter := structservices.NewFunctionInterpreter()
	discover := structservices.NewDiscover(file, classifier, interfaceInterpreter, functionInterpreter)

	files, err := discover.Project("docs/examples/medium")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		f.Print("")
	}
}
