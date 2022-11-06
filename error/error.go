package error_type

import "errors"

func ErrorFun(i int) (int, error) {
	if i < 0 {
		return 0, errors.New("error")
	}
	return i, nil
}
