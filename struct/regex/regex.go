package regex

import "regexp"

type Regex struct {
	Case    *regexp.Regexp
	Ignores []*regexp.Regexp
}

func NewRegex(_case *regexp.Regexp, ignores ...*regexp.Regexp) *Regex {
	return &Regex{
		Case:    _case,
		Ignores: ignores,
	}
}

func (r *Regex) FindString(s string) string {
	for _, ignore := range r.Ignores {
		find := ignore.FindString(s)
		if find != "" {
			return ""
		}
	}

	return r.Case.FindString(s)
}

func (r *Regex) FindAllString(s string, n int) []string {
	for _, ignore := range r.Ignores {
		find := ignore.FindAllString(s, n)
		if len(find) > 0 {
			return nil
		}
	}

	return r.Case.FindAllString(s, n)
}
