package structmodels

import (
	"fmt"
	"strings"
)

type Import struct {
	Alias string
	Full  string
}

func NewImport(full string, alias ...string) Import {
	i := Import{
		Full:  full,
		Alias: getAlias(full),
	}
	if len(alias) > 0 {
		i.Alias = alias[0]
	}
	return i
}

func (i *Import) Import() string {
	values := strings.Split(i.Full, "/")

	if values[len(values)-1] == i.Alias {
		return fmt.Sprintf(`"%s"`, i.Full)
	}

	return fmt.Sprintf(`%s "%s"`, i.Alias, i.Full)
}

func (i *Import) String() string {
	return fmt.Sprintf(`import "%s"`, i.Full)
}

type Imports []Import

func (i Imports) String() string {
	if len(i) == 0 {
		return ""
	}
	if len(i) == 1 {
		return i[0].String()
	}
	s := "import ("
	for _, _import := range i {
		s += "\n\t" + fmt.Sprintf(`"%s"`, _import.Full)
	}
	s += "\n)\n"
	return s
}

func getAlias(full string) string {
	words := strings.Split(full, "/")
	if len(words) == 1 {
		return words[0]
	}

	return words[len(words)-1]
}
