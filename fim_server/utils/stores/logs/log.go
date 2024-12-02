package logs

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func log(color string, s ...interface{}) string {
	str := isSprintf(s...)
	t := logColor(fmt.Sprintf("[%v]", time.Now().Format(times)), "#ffffff")
	fmt.Println(t + getLine(line) + logColor(isFieldColor(str), color))
	return str
}

func Info(s ...interface{}) string {
	return log("#80ffff", s...)
}

func Error(s ...interface{}) error {
	return errors.New(log("#F75464", s...))
}

func Fatal(s ...interface{}) {
	log("#F75464", s...)
	os.Exit(1)
}
