package logs

import (
	"errors"
	"fmt"
	"os"
)

var c = struct {
	Line  int
	Info  string
	Error string
	Time  string
}{
	Line:  1,
	Time:  "15:04:05",
	Info:  "#80ffff",
	Error: "#F75464",
}

func Info(s ...interface{}) string {
	return log(getLine(), c.Info, s...).Info()
}
func InfoF(sf string, s ...interface{}) string {
	return log(getLine(), c.Info, fmt.Sprintf(sf, s...)).Info()
}
func Error(s ...interface{}) error {
	return errors.New(log(getLine(), c.Error, s...).Error())
}
func Fatal(s ...interface{}) {
	log(getLine(), c.Error, s...).Error()
	os.Exit(0)
}
