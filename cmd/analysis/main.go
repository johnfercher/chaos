package main

import (
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"github.com/johnfercher/chaos/struct/structservices"
	"log"
)

func main() {
	file := structservices.NewFile()
	classifier := structservices.NewFileClassifier()
	discover := structservices.NewDiscover(file, classifier)
	interfaceInterpreter := structservices.NewInterfaceInterpreter()

	files, err := discover.Project("docs/examples/medium")
	if err != nil {
		log.Fatal(err)
	}

	var interfaces []*structmodels.Interface
	for _, f := range files {
		f.Print("")
		_interface := interfaceInterpreter.Interpret(f.Content)
		interfaces = append(interfaces, _interface...)
	}
}
