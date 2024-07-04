package structmodels

import (
	"fmt"
	"github.com/johnfercher/chaos/struct/structcore/structconsts/content"
	"github.com/johnfercher/chaos/struct/structcore/structconsts/file"
)

type File struct {
	Name        string
	Path        string
	FullPath    string
	Type        file.Type
	ContentType content.Type
	Package     string
	Content     string
}

func (f *File) Print(identation string) {
	fmt.Printf("%s%s, %s, %s, %s, %s\n", identation, f.Type, f.ContentType, f.Package, f.Name, f.Path)
	/*if f.Type == file.Dir {
		fmt.Printf("%s/%s\n", identation, f.Name)
	} else {
		fmt.Printf("%s%s\n", identation, f.Name)
	}*/
}
