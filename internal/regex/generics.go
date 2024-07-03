package regex

import (
	"regexp"
)

var openParantesis = regexp.MustCompile(`\(`)
var closeParantesis = regexp.MustCompile(`\)`)
var openBrackets = regexp.MustCompile(`\{`)
var closeBrackets = regexp.MustCompile(`\}`)
