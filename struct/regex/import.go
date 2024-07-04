package regex

import (
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"regexp"
	"strings"
)

var (
	singleLineImports = NewRegex(regexp.MustCompile(`import\s.+`))
	multiLineImports  = NewRegex(regexp.MustCompile(`import\s\(`))
)

func GetImports(file string) []structmodels.Import {
	singleLineFullImports := singleLineImports.FindAllString(file, -1)
	var imports []structmodels.Import

	for _, singleLineFullImport := range singleLineFullImports {
		if singleLineFullImport == "import (" {
			continue
		}
		_import := strings.ReplaceAll(singleLineFullImport, `import "`, "")
		_import = strings.ReplaceAll(_import, `"`, "")
		imports = append(imports, getImport(_import))
	}

	multipleImports := getMultipleImports(file)
	imports = append(imports, multipleImports...)

	return imports
}

func getMultipleImports(file string) []structmodels.Import {
	scope := GetMultiLineScope(file, multiLineImports, closeParantesis)
	lines := strings.Split(scope, "\n")

	_imports := []structmodels.Import{}
	for i := 1; i < len(lines)-2; i++ {
		line := lines[i]
		_imports = append(_imports, getImport(line))
	}

	return _imports
}

func getImport(line string) structmodels.Import {
	line = strings.ReplaceAll(line, "\t", "")
	line = strings.ReplaceAll(line, `"`, "")
	values := strings.Split(line, " ")

	if len(values) == 1 {
		return structmodels.NewImport(values[0])
	}
	return structmodels.NewImport(values[1], values[0])
}
