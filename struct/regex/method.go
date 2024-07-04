package regex

import (
	"fmt"
	"github.com/johnfercher/chaos/struct/structcore/structmodels"
	"regexp"
	"strings"
)

var methodName = regexp.MustCompile(`\w+\[?\w+(\s?\w+)?\]?\(`)

func GetMethod(method string) structmodels.Method {
	m := structmodels.Method{}

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
		m.Parameters = append(m.Parameters, structmodels.NewParameter(values[0], values[1]))
	}

	if returnScope == "" {
		return m
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

func GetParameter(parameterString string) structmodels.Parameter {
	parameter := strings.Split(parameterString, " ")

	if len(parameter) == 1 {
		return structmodels.NewParameter("", parameter[0])
	}

	return structmodels.NewParameter(parameter[0], parameter[1])
}
