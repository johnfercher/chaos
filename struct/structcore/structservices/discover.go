package structservices

import "github.com/johnfercher/chaos/struct/structcore/structmodels"

type Discover interface {
	Project(path string) ([]structmodels.File, error)
}
