package logs

import (
	"errors"
	"fmt"
	"time"
)

func Info(s ...interface{}) string {
	str := isSprintf(s...)
	t := logColor(fmt.Sprintf("[%v]", time.Now().Format(times)), "#ffffff")
	fmt.Println(t + getLine(line) + logColor(isFieldColor(str), "#80ffff"))
	return str
}

func Error(s ...interface{}) error {
	str := isSprintf(s...)
	t := logColor(fmt.Sprintf("[%v]", time.Now().Format(times)), "#ffffff")
	fmt.Println(t + getLine(line) + logColor(isFieldColor(str), "#F75464"))

	return errors.New(str)
}
