package error_messages

import "strconv"

func DataBetween(col string, min int, max int) string{
	return col + " should between" + strconv.Itoa(min) + " and " + strconv.Itoa(max)
}