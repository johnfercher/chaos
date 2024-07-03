package regex

import (
	"github.com/johnfercher/chaos/internal/model"
	"regexp"
	"strings"
)

/*
Add(ctx context.Context, id string) error
Add[T](ctx context.Context, id string) error
Add[T any](ctx context.Context, id string) error
*/
var methodName = regexp.MustCompile(`\w+\[?\w+(\s?\w+)?\]?\(`)

var builtinParameter = regexp.MustCompile(`\w+\s\w+(,|\))`)
var packageParameter = regexp.MustCompile(`\w+\s\w+\.\w+(,|\))`)

var builtinPointerParameter = regexp.MustCompile(`\w+\s\*\w+(,|\))`)
var packagePointerParameter = regexp.MustCompile(`\w+\s\*\w+\.\w+(,|\))`)

var builtinVarArgParameter = regexp.MustCompile(`\w+\s(\.\.\.)\w+(,|\))`)
var packageVarArgParameter = regexp.MustCompile(`\w+\s(\.\.\.)\w+\.\w+(,|\))`)

var builtinVarArgPointerParameter = regexp.MustCompile(`\w+\s(\.\.\.)\*\w+(,|\))`)
var packageVarArgPointerParameter = regexp.MustCompile(`\w+\s(\.\.\.)\*\w+\.\w+(,|\))`)

var builtinArrayParameter = regexp.MustCompile(`\w+\s\[\]\w+(,|\))`)
var packageArrayParameter = regexp.MustCompile(`\w+\s\[\]\w+\.\w+(,|\))`)

var builtinArrayPointerParameter = regexp.MustCompile(`\w+\s\[\]\*\w+(,|\))`)
var packageArrayPointerParameter = regexp.MustCompile(`\w+\s\[\]\*\w+\.\w+(,|\))`)

var trueBuiltinMapParameter = regexp.MustCompile(`\w+\s\w+\[\w+\]\w+(,|\))`)
var keyPackageMapBuiltinValueParameter = regexp.MustCompile(`\w+\s\w+\[\w+\.\w+\]\w+(,|\))`)
var keyBuiltinMapPackageValueParameter = regexp.MustCompile(`\w+\s+\w+\[\w+\]\w+\.\w+(,|\))`)
var truePackageMapParameter = regexp.MustCompile(`\w+\s\w+\[\w+\.\w+\]\w+\.\w+(,|\))`)

func GetMethod(method string) model.Method {
	m := model.Method{}

	name := methodName.FindString(method)
	name = strings.ReplaceAll(name, "(", "")
	m.Name = name

	method = strings.ReplaceAll(method, name, "")

	parametersScope := GetSingleLineScope(method, "(", ")")
	parametersScope = strings.ReplaceAll(parametersScope, "(", "")
	parametersScope = strings.ReplaceAll(parametersScope, ")", "")
	parameters := strings.Split(parametersScope, ",")
	for i := 0; i < len(parameters); i++ {
		parameters[i] += ", "
	}

	for _, parameter := range parameters {
		m.Parameters = append(m.Parameters, GetBuiltinParameter(parameter)...)
		m.Parameters = append(m.Parameters, GetPackageParameter(parameter)...)

		m.Parameters = append(m.Parameters, GetBuiltinPointerParameter(parameter)...)
		m.Parameters = append(m.Parameters, GetPackagePointerParameter(parameter)...)

		m.Parameters = append(m.Parameters, GetBuiltinVarArgParameter(parameter)...)
		m.Parameters = append(m.Parameters, GetPackageVarArgParameter(parameter)...)

		m.Parameters = append(m.Parameters, GetBuiltinVarArgPointerParameter(parameter)...)
		m.Parameters = append(m.Parameters, GetPackageVarArgPointerParameter(parameter)...)

		m.Parameters = append(m.Parameters, GetBuiltinArrayParameter(parameter)...)
		m.Parameters = append(m.Parameters, GetPackageArrayParameter(parameter)...)

		m.Parameters = append(m.Parameters, GetBuiltinArrayPointerParameter(parameter)...)
		m.Parameters = append(m.Parameters, GetPackageArrayPointerParameter(parameter)...)

		m.Parameters = append(m.Parameters, GetTrueBuiltinMapParameter(parameter)...)
		m.Parameters = append(m.Parameters, GetKeyPackageMapBuiltinValueParameter(parameter)...)
		m.Parameters = append(m.Parameters, GetKeyBuiltinMapPackageValueParameter(parameter)...)
		m.Parameters = append(m.Parameters, GetTruePackageMapParameter(parameter)...)
	}

	return m
}

func GetBuiltinParameter(method string) []model.Parameter {
	_strings := builtinParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetPackageParameter(method string) []model.Parameter {
	_strings := packageParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetBuiltinPointerParameter(method string) []model.Parameter {
	_strings := builtinPointerParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetPackagePointerParameter(method string) []model.Parameter {
	_strings := packagePointerParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetBuiltinVarArgParameter(method string) []model.Parameter {
	_strings := builtinVarArgParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetPackageVarArgParameter(method string) []model.Parameter {
	_strings := packageVarArgParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetBuiltinVarArgPointerParameter(method string) []model.Parameter {
	_strings := builtinVarArgPointerParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetPackageVarArgPointerParameter(method string) []model.Parameter {
	_strings := packageVarArgPointerParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetBuiltinArrayParameter(method string) []model.Parameter {
	_strings := builtinArrayParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetPackageArrayParameter(method string) []model.Parameter {
	_strings := packageArrayParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetBuiltinArrayPointerParameter(method string) []model.Parameter {
	_strings := builtinArrayPointerParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetPackageArrayPointerParameter(method string) []model.Parameter {
	_strings := packageArrayPointerParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetTrueBuiltinMapParameter(method string) []model.Parameter {
	_strings := trueBuiltinMapParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetKeyPackageMapBuiltinValueParameter(method string) []model.Parameter {
	_strings := keyPackageMapBuiltinValueParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetKeyBuiltinMapPackageValueParameter(method string) []model.Parameter {
	_strings := keyBuiltinMapPackageValueParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func GetTruePackageMapParameter(method string) []model.Parameter {
	_strings := truePackageMapParameter.FindAllString(method, -1)
	var parameters []model.Parameter

	for _, s := range _strings {
		parameters = append(parameters, buildParameter(s))
	}

	return parameters
}

func buildParameter(s string) model.Parameter {
	values := strings.Split(s, " ")
	name := values[0]
	t := values[1]
	t = strings.ReplaceAll(t, ",", "")
	t = strings.ReplaceAll(t, ")", "")
	return model.Parameter{
		Name: name,
		Type: t,
	}
}
