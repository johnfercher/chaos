package main

import (
	"bytes"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"github.com/johnfercher/chaos/struct/structservices"
	"log"
	"os"
	"strings"
)

func main() {
	file := structservices.NewFile()
	classifier := structservices.NewFileClassifier()
	interfaceInterpreter := structservices.NewInterfaceInterpreter()
	functionInterpreter := structservices.NewFunctionInterpreter()
	discover := structservices.NewDiscover(file, classifier, interfaceInterpreter, functionInterpreter)

	path := "docs/examples/medium"
	files, err := discover.Project(path)
	if err != nil {
		log.Fatal(err)
	}

	var mod structmodels.File
	for _, f := range files {
		if f.Name == "go.mod" {
			mod = f
		}
	}

	goModFirstLine := strings.Split(mod.Content, "\n")[0]
	goModFirstLine = strings.ReplaceAll(goModFirstLine, "module ", "")

	for i := 0; i < len(files); i++ {
		//files[i].Path = strings.ReplaceAll(files[i].Path, path+"/", "")
		for j := 0; j < len(files[i].Imports); j++ {
			files[i].Imports[j].Full = strings.ReplaceAll(files[i].Imports[j].Full, goModFirstLine+"/", "")
		}
	}

	links := MergeN2(files)

	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	nodes := make(map[string]*cgraph.Node)

	for key, _ := range links {
		n, err := graph.CreateNode(key)
		if err != nil {
			log.Fatal(err)
		}
		nodes[key] = n
	}

	/*e, err := graph.CreateEdge("e", n, m)
	if err != nil {
		log.Fatal(err)
	}*/
	//e.SetLabel("e")
	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("test.gv", buf.Bytes(), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.ReadFile("test.gv")
	if err != nil {
		log.Fatal(err)
	}

	if err := g.RenderFilename(graph, graphviz.PNG, "graph.png"); err != nil {
		log.Fatal(err)
	}

	/*for _, f := range files {
		f.Print()
	}*/
}

func MergeN2(files []structmodels.File) map[string]map[string]bool {
	m := make(map[string]map[string]bool)
	for _, file := range files {
		m2 := make(map[string]bool)
		for _, f2 := range file.Imports {
			m2[f2.Full] = true
		}
		m2Active, ok := m[file.Path]
		if ok {
			m3 := MergeN1(m2Active, m2)
			m[file.Path] = m3
		} else {
			m[file.Path] = m2
		}

	}
	return m
}

func MergeN1(a map[string]bool, b map[string]bool) map[string]bool {
	for key, v := range b {
		a[key] = v
	}

	return a
}
