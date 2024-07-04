package regex

import (
	"regexp"
	"strings"
)

var (
	singleLineImports = regexp.MustCompile(`import\s.+`)
	multiLineImports  = regexp.MustCompile(`import\s\(`)
)

func GetImports(file string) []string {
	singleLineFullImports := singleLineImports.FindAllString(file, -1)
	var imports []string

	for _, singleLineFullImport := range singleLineFullImports {
		if singleLineFullImport == "import (" {
			continue
		}
		_import := strings.ReplaceAll(singleLineFullImport, `import "`, "")
		_import = strings.ReplaceAll(_import, `"`, "")
		imports = append(imports, _import)
	}

	multipleImports := getMultipleImports(file)
	imports = append(imports, multipleImports...)

	return imports
}

func getMultipleImports(file string) []string {
	scope := GetMultiLineScope(file, multiLineImports, closeParantesis)
	lines := strings.Split(scope, "\n")

	_strings := []string{}
	for i := 1; i < len(lines)-2; i++ {
		line := lines[i]
		line = strings.ReplaceAll(line, "\t", "")
		line = strings.ReplaceAll(line, `"`, "")
		_strings = append(_strings, line)
	}

	return _strings
}
