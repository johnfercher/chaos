package regex

import (
	"regexp"
)

var (
	openParantesis  = regexp.MustCompile(`\(`)
	closeParantesis = regexp.MustCompile(`\)`)
	openBrackets    = regexp.MustCompile(`\{`)
	closeBrackets   = regexp.MustCompile(`\}`)
)
