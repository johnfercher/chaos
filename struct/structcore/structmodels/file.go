package structmodels

import (
	"fmt"
	"github.com/johnfercher/chaos/struct/structcore/structconsts/content"
	"github.com/johnfercher/chaos/struct/structcore/structconsts/file"
)

type File struct {
	Name        string
	Path        string
	Type        file.Type
	ContentType content.Type
	Package     string
	Imports     Imports
	Content     string
	Interfaces  []*Interface
	Functions   []Method
}

func (f *File) Print() {
	fmt.Printf("%s, %s, %s, %s, %s\n", f.Type, f.ContentType, f.Package, f.Name, f.Path)
	fmt.Println(f.Imports.String())
}
