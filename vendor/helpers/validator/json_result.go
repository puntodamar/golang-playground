package validator

import (
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