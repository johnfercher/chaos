package regex

import (
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"regexp"
	"strings"
)

var (
	singleLineImports = regexp.MustCompile(`import\s.+`)
	multiLineImports  = regexp.MustCompile(`import\s\(`)
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
		imports = append(imports, structmodels.NewImport(_import))
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
		line = strings.ReplaceAll(line, "\t", "")
		line = strings.ReplaceAll(line, `"`, "")
		_imports = append(_imports, structmodels.NewImport(line))
	}

	return _imports
}
