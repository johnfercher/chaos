package main

import (
	"fmt"
	"github.com/johnfercher/chaos/struct/structservices"
	"log"
)

func main() {
	file := structservices.NewFile()
	classifier := structservices.NewFileClassifier()
	discover := structservices.NewDiscover(file, classifier)
	ok, err := discover.Project("docs/examples/medium")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ok)
}
