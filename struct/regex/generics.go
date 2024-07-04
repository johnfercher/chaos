package regex

import (
	"regexp"
)

var (
	openParantesis  = NewRegex(regexp.MustCompile(`\(`))
	closeParantesis = NewRegex(regexp.MustCompile(`\)`))
	openBrackets    = NewRegex(regexp.MustCompile(`\{`))
	closeBrackets   = NewRegex(regexp.MustCompile(`\}`), regexp.MustCompile(`interface\{\}`))
)
