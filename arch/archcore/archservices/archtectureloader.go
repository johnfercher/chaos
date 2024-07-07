package archservices

import "github.com/johnfercher/chaos/arch/archcore/archmodels"

type Loader interface {
	LoadArchitecture(file string) (*archmodels.Node, error)
}
