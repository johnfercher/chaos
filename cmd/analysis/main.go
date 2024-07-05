package main

import (
	"fmt"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"github.com/johnfercher/chaos/struct/structservices"
	"log"
)

func main() {
	file := structservices.NewFile()
	classifier := structservices.NewFileClassifier()
	discover := structservices.NewDiscover(file, classifier)
	interfaceInterpreter := structservices.NewInterfaceInterpreter()
	functionInterpreter := structservices.NewFunctionInterpreter()

	files, err := discover.Project("docs/examples/medium")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		f.Print("")
	}

	var interfaces []*structmodels.Interface
	for _, f := range files {
		_interface := interfaceInterpreter.Interpret(f.Content)
		interfaces = append(interfaces, _interface...)
	}

	for _, _interface := range interfaces {
		fmt.Println(_interface.String())
	}

	var functions []structmodels.Method
	for _, f := range files {
		function := functionInterpreter.Interpret(f.Content)
		functions = append(functions, function...)
	}

	for _, function := range functions {
		fmt.Println(function.String())
	}
}
