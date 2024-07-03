package regex

import (
	"fmt"
	"github.com/johnfercher/chaos/internal/model"
	"regexp"
	"strings"
)

var PackageRegex = regexp.MustCompile(`package\s\w+`)
var SingleLineImports = regexp.MustCompile(`import\s.+`)
var MultiLineImports = regexp.MustCompile(`import\s\(`)
var CloseParantesis = regexp.MustCompile(`\)`)
var CloseBrackets = regexp.MustCompile(`\}`)
var InterfaceName = regexp.MustCompile(`type\s.+interface\s+{`)

func GetScope(file string, begin *regexp.Regexp, end *regexp.Regexp) string {
	lines := strings.Split(file, "\n")
	beginID := -1
	for id, line := range lines {
		s := begin.FindString(line)
		if s != "" {
			beginID = id
			break
		}
	}

	if beginID < 0 {
		return ""
	}

	_strings := []string{}
	for i := beginID; i < len(lines); i++ {
		line := lines[i]
		_strings = append(_strings, line+"\n")
		s := end.FindString(line)
		if s != "" {
			break
		}
	}

	return strings.Join(_strings, "")
}

func GetPackageName(file string) string {
	s := PackageRegex.FindString(file)
	return strings.ReplaceAll(s, "package ", "")
}

func GetImports(file string) []string {
	singleLineFullImports := SingleLineImports.FindAllString(file, -1)
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

func GetInterfaces(file string) []*model.Interface {
	fullInterfaces := InterfaceName.FindAllString(file, -1)
	var interfaces []*model.Interface

	for _, fullInterface := range fullInterfaces {
		_interface := strings.ReplaceAll(fullInterface, "type ", "")
		_interface = strings.ReplaceAll(_interface, " interface {", "")
		interfaces = append(interfaces, &model.Interface{
			Name: _interface,
		})
	}

	for i, _interface := range interfaces {
		methods := getInterfaceMethods(file, _interface.Name)
		interfaces[i].Methods = methods
	}

	return interfaces
}

func getInterfaceMethods(file string, name string) []string {
	pattern := fmt.Sprintf(`type\s%s\sinterface\s+{`, name)
	begin := regexp.MustCompile(pattern)

	scope := GetScope(file, begin, CloseBrackets)
	fmt.Println(scope)

	lines := strings.Split(scope, "\n")

	_strings := []string{}
	for i := 1; i < len(lines)-2; i++ {
		line := lines[i]
		line = strings.ReplaceAll(line, "\t", "")
		_strings = append(_strings, line)
	}

	return _strings
}

func getMultipleImports(file string) []string {
	scope := GetScope(file, MultiLineImports, CloseParantesis)
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
