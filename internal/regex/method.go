package regex

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/johnfercher/chaos/internal/core/models"
)

var methodName = regexp.MustCompile(`\w+\[?\w+(\s?\w+)?\]?\(`)

func GetMethod(method string) models.Method {
	m := models.Method{}

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
		m.Parameters = append(m.Parameters, models.NewParameter(values[0], values[1]))
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

func GetParameter(parameterString string) models.Parameter {
	parameter := strings.Split(parameterString, " ")

	if len(parameter) == 1 {
		return models.NewParameter("", parameter[0])
	}

	return models.NewParameter(parameter[0], parameter[1])
}
