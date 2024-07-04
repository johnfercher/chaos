package regex

import (
	"regexp"
	"strings"
)

var packageRegex = NewRegex(regexp.MustCompile(`package\s\w+`))

func GetPackageName(file string) string {
	s := packageRegex.FindString(file)
	return strings.ReplaceAll(s, "package ", "")
}
