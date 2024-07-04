package structmodels

import (
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
