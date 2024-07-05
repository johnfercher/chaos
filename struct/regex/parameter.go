package regex

import (
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"strings"
)

func GetParameters(line string) []structmodels.Parameter {
	if line == "" {
		return nil
	}

	var p []structmodels.Parameter
	parameters := strings.Split(line, ",")
	for i := 0; i < len(parameters); i++ {
		if i > 0 {
			parameters[i] = strings.Replace(parameters[i], " ", "", 1)
		}
	}

	lastType := ""
	for i := len(parameters) - 1; i >= 0; i-- {
		values := strings.Split(parameters[i], " ")
		if len(values) > 1 {
			lastType = values[1]
		} else {
			parameters[i] += " " + lastType
		}
	}

	for _, parameter := range parameters {
		values := strings.Split(parameter, " ")
		p = append(p, structmodels.NewParameter(values[0], values[1]))
	}

	return p
}
