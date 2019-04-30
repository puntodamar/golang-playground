package error_formats

type ErrorFormat struct{
	Field 		string 			`json:"field"`
	Messages 	[]interface{}	`json:"messages"`
}