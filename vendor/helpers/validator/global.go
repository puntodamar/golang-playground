package validator

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)

	//// after
	//govalidator.CustomTypeTagMap.Set("customByteArrayValidator", CustomTypeValidator(func(i interface{}, o interface{}) bool {
	//	// ...
	//}))

	// Add your own struct validation tags
	govalidator.TagMap["exists"] = govalidator.Validator(func(str string) bool {
		return str == "duck"
	})
}
