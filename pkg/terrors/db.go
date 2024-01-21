package terrors

import "strings"

func IsNotFoundErr(err error) bool {
	return strings.Contains(err.Error(), "no rows")
}

func NewDbErr(err error) Error {
	if IsNotFoundErr(err) {
		return NewNotFoundError("not found: ", nil)
	}

	return NewPrivateError(err.Error())
}
