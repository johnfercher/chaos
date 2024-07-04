package structmodels

import "strings"

type Import struct {
	Alias string
	Full  string
}

func (i *Import) Import() string {
	return i.Full
}

func NewImport(full string) Import {
	return Import{
		Full:  full,
		Alias: getAlias(full),
	}
}

func getAlias(full string) string {
	words := strings.Split(full, "/")
	if len(words) == 1 {
		return words[0]
	}

	return words[len(words)-1]
}
