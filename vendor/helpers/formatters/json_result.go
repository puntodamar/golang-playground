package formatters

import (
	"fmt"
	"strings"
)

type SuccessMessage struct {
	Message string		`json:"message"`
	Data 	interface{}	`json:"data"`
}

type JsonFormatter struct {
	stringErrors 	string
	mapErrors 		map[string]string
}

type ErrorField struct {
	Field 	string
	Errors 	[]string
}

func (*JsonFormatter) FormatJsonErrorFromString(err string) map[string]string{

	errorList 		:= strings.Split(err, ";")
	jsonErrorMap 	:= map[string]string{}

	for _, field := range errorList{
		split := strings.Split(field,":")
		jsonErrorMap[split[0]] = split[1]
	}

	return jsonErrorMap
}

func (*JsonFormatter) FormatJsonErrorFromMap(err map[string]string) map[string]map[string]string{

	jsonErrorMap 	:= map[string]string{}

	fmt.Println(err)
	for field, msg := range err{
		jsonErrorMap[field] = msg

	}

	return map[string]map[string]string{"errors" : jsonErrorMap}
}

func (*JsonFormatter) FormatJsonSuccess(data string, method string, value interface{}) SuccessMessage{

	msg := SuccessMessage{
		Message	: data + " has been " + method + "d successfully",
		Data	: value,
	}

	return msg
}

func (*JsonFormatter) JoinValidationErrors(original map[string]string, data map[string]string) []ErrorField {

	var final []ErrorField

	for key, v1 := range original {

		e  := ErrorField{Field:key, Errors: []string{v1}}

		if v2, exists := data[v1]; exists {
			e.Errors = append(e.Errors, v2)
		}

		final = append(final, e)
	}


	return final
}