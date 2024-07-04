package regex

import (
	"strings"
)

func GetMultiLineScope(file string, begin *Regex, end *Regex) string {
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

func GetSingleLineScope(line string, openCharacter string, closeCharacter string) string {
	s := ""

	open := []rune(openCharacter)
	close := []rune(closeCharacter)
	found := false
	for _, character := range line {
		if character == open[0] {
			found = true
		}
		if found {
			s += string(character)
		}
		if character == close[0] {
			break
		}
	}

	return s
}
