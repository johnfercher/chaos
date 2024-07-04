package main

import (
	"github.com/johnfercher/chaos/struct/structservices"
	"log"
)

func main() {
	file := structservices.NewFile()
	classifier := structservices.NewFileClassifier()
	discover := structservices.NewDiscover(file, classifier)
	packages, err := discover.Project("docs/examples/medium")
	if err != nil {
		log.Fatal(err)
	}

	for _, pkg := range packages {
		pkg.Print("")
	}
}
