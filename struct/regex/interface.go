package regex

import (
	"fmt"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"regexp"
	"strings"
)

var interfaceName = NewRegex(regexp.MustCompile(`type\s.+interface\s+{`))

func GetInterfaces(file string) []*structmodels.Interface {
	fullInterfaces := interfaceName.FindAllString(file, -1)
	var interfaces []*structmodels.Interface

	for _, fullInterface := range fullInterfaces {
		_interface := strings.ReplaceAll(fullInterface, "type ", "")
		_interface = strings.ReplaceAll(_interface, " interface {", "")
		interfaces = append(interfaces, &structmodels.Interface{
			Name: _interface,
		})
	}

	for i, _interface := range interfaces {
		methods := getInterfaceMethods(file, _interface.Name)
		interfaces[i].Methods = methods
	}

	return interfaces
}

func getInterfaceMethods(file string, name string) []structmodels.Method {
	pattern := fmt.Sprintf(`type\s%s\sinterface\s+{`, name)
	begin := NewRegex(regexp.MustCompile(pattern))

	scope := GetMultiLineScope(file, begin, closeBrackets)

	lines := strings.Split(scope, "\n")

	methods := []structmodels.Method{}
	for i := 1; i < len(lines)-2; i++ {
		line := lines[i]
		line = strings.ReplaceAll(line, "\t", "")
		m := GetMethod(line)
		m.Source = name
		methods = append(methods, m)
	}

	return methods
}
