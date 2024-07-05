package regex

import (
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"regexp"
	"strings"
)

var functionName = NewRegex(regexp.MustCompile(`func\s\w+\[?\w+(\s?\w+)?\]?\(`))

func GetFunctions(file string) []structmodels.Method {
	var methods []structmodels.Method

	lines := strings.Split(file, "\n")
	for _, line := range lines {
		s := functionName.FindString(line)
		if s != "" {
			line = strings.ReplaceAll(line, "func ", "")
			line = strings.ReplaceAll(line, " {", "")
			m := GetMethod(line)
			methods = append(methods, m)
		}
	}
	return methods
}
