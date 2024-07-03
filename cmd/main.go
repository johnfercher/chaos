package main

import (
	"fmt"
	"github.com/johnfercher/chaos/internal"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	reader := internal.NewFileReader(dir)

	file, err := reader.Load("/internal/example/interface.go")
	if err != nil {
		log.Fatal(err)
	}

	interpreter := internal.NewInterfaceInterpreter()
	interfaces := interpreter.Interpret(file)

	fmt.Println(interfaces)

	decoratorTemplate, err := reader.Load("/internal/template/chaos/decorator.txt")
	if err != nil {
		log.Fatal(err)
	}

	methodTemplate, err := reader.Load("/internal/template/chaos/method.txt")
	if err != nil {
		log.Fatal(err)
	}

	decorator := internal.NewDecoratorGenerator(decoratorTemplate, methodTemplate)
	for _, _interface := range interfaces {
		s := decorator.Generate(_interface)
		fmt.Println(s)
	}
}
