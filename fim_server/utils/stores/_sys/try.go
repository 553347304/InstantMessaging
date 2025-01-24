package _sys

import (
	"errors"
	"fmt"
)

func Try(errFuncList ...func() error) error {
	for i, errFunc := range errFuncList {
		err := errFunc()
		if err != nil {
			return errors.New(fmt.Sprint("函数", i, "->", err.Error()))
		}
	}
	return nil
}
