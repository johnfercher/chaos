package structmodels

import (
	"fmt"
	"github.com/johnfercher/chaos/struct/structcore/structconsts/content"
	"github.com/johnfercher/chaos/struct/structcore/structconsts/file"
)

type Package struct {
	Name        string
	Path        string
	Type        file.Type
	ContentType content.Type
	Files       []File
	Packages    []Package
}

func (f *Package) Print(identation string) {
	fmt.Println("")
	fmt.Printf("%s%s, %s, %s, %s\n", identation, f.Type, f.ContentType, f.Name, f.Path)
	/*if f.Type == file.Dir {
		fmt.Printf("%s/%s\n", identation, f.Name)
	} else {
		fmt.Printf("%s%s\n", identation, f.Name)
	}*/
}
