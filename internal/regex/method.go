package regex

import (
	"fmt"
	"github.com/johnfercher/chaos/internal/model"
	"regexp"
	"strings"
)

var methodName = regexp.MustCompile(`\w+\[?\w+(\s?\w+)?\]?\(`)

func GetMethod(method string) model.Method {
	m := model.Method{}

	name := methodName.FindString(method)
	name = strings.ReplaceAll(name, "(", "")
	m.Name = name

	method = strings.ReplaceAll(method, name, "")

	parametersScope := GetSingleLineScope(method, "(", ")")
	returnScope := strings.ReplaceAll(method, parametersScope, "")
	parametersScope = strings.ReplaceAll(parametersScope, "(", "")
	parametersScope = strings.ReplaceAll(parametersScope, ")", "")
	parametersScope = strings.ReplaceAll(parametersScope, ", ", ",")
	parameters := strings.Split(parametersScope, ",")

	for _, parameter := range parameters {
		values := strings.Split(parameter, " ")
		m.Parameters = append(m.Parameters, model.Parameter{
			Name: values[0],
			Type: values[1],
		})
	}

	returnScope = strings.Replace(returnScope, " ", "", 1)
	returnScope = strings.ReplaceAll(returnScope, "(", "")
	returnScope = strings.ReplaceAll(returnScope, ")", "")
	returns := strings.Split(returnScope, ",")
	for i := 0; i < len(returns); i++ {
		r := fmt.Sprintf("%c", returns[i][0])
		if r == " " {
			returns[i] = strings.Replace(returns[i], " ", "", 1)
		}
	}

	for _, _return := range returns {
		m.Returns = append(m.Returns, GetParameter(_return))
	}

	return m
}

func GetParameter(parameterString string) model.Parameter {
	parameter := strings.Split(parameterString, " ")
	p := model.Parameter{}

	if len(parameter) == 1 {
		p.Type = parameter[0]
	} else {
		p.Name = parameter[0]
		p.Type = parameter[1]
	}
	return p
}
