package archservices

import (
	"fmt"
	"github.com/johnfercher/chaos/arch/archcore/archmodels"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"strings"
)

type ImportsDiscover struct {
}

func NewImportsDiscover() *ImportsDiscover {
	return &ImportsDiscover{}
}

func (i *ImportsDiscover) Find(files []structmodels.File) archmodels.Project {
	var mod structmodels.File
	for _, f := range files {
		if f.Name == "go.mod" {
			mod = f
		}
	}

	path := mod.Path

	goModFirstLine := strings.Split(mod.Content, "\n")[0]
	goModFirstLine = strings.ReplaceAll(goModFirstLine, "module ", "")

	for i := 0; i < len(files); i++ {
		files[i].Path = strings.ReplaceAll(files[i].Path, path+"/", "")
		for j := 0; j < len(files[i].Imports); j++ {
			files[i].Imports[j].Full = strings.ReplaceAll(files[i].Imports[j].Full, goModFirstLine+"/", "")
		}
	}

	packages := MergeN2(files)

	for key, imports := range packages {
		fmt.Println(key)
		for key2, _ := range imports {
			fmt.Printf("\t%s\n", key2)
		}
	}

	return packages
}

func MergeN2(files []structmodels.File) map[string]map[string]bool {
	m := make(map[string]map[string]bool)
	for _, file := range files {
		m2 := make(map[string]bool)
		for _, f2 := range file.Imports {
			m2[f2.Alias] = true
		}
		values := strings.Split(file.Path, "/")
		lastValue := values[len(values)-1]
		m2Active, ok := m[lastValue]
		if ok {
			m3 := MergeN1(m2Active, m2)
			m[lastValue] = m3
		} else {
			m[lastValue] = m2
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
