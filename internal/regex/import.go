package regex

import (
	"github.com/johnfercher/chaos/internal/core/models"
	"regexp"
	"strings"
)

var (
	singleLineImports = regexp.MustCompile(`import\s.+`)
	multiLineImports  = regexp.MustCompile(`import\s\(`)
)

func GetImports(file string) []models.Import {
	singleLineFullImports := singleLineImports.FindAllString(file, -1)
	var imports []models.Import

	for _, singleLineFullImport := range singleLineFullImports {
		if singleLineFullImport == "import (" {
			continue
		}
		_import := strings.ReplaceAll(singleLineFullImport, `import "`, "")
		_import = strings.ReplaceAll(_import, `"`, "")
		imports = append(imports, models.NewImport(_import))
	}

	multipleImports := getMultipleImports(file)
	imports = append(imports, multipleImports...)

	return imports
}

func getMultipleImports(file string) []models.Import {
	scope := GetMultiLineScope(file, multiLineImports, closeParantesis)
	lines := strings.Split(scope, "\n")

	_imports := []models.Import{}
	for i := 1; i < len(lines)-2; i++ {
		line := lines[i]
		line = strings.ReplaceAll(line, "\t", "")
		line = strings.ReplaceAll(line, `"`, "")
		_imports = append(_imports, models.NewImport(line))
	}

	return _imports
}
