package archservices

import (
	"github.com/johnfercher/chaos/arch/archcore/archmodels"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
)

type ImportstDiscover interface {
	Find(files []structmodels.File) archmodels.Project
}
