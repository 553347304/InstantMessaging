package logs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var c = struct {
	Info  string
	Warn  string
	Error string
	Time  string
}{
	Time:  "15:04:05",
	Info:  "#80ffff",
	Warn:  "#DFAD49",
	Error: "#F75464",
}

func Info(s ...interface{}) string {
	fmt.Println(l.time(), l.line()[0], l.color(l.ip(l.string(s...)), c.Info))
	return l.string(s...)
}
func InfoF(sf string, s ...interface{}) {
	fmt.Println(l.time(), l.line()[0], l.color(l.ip(fmt.Sprintf(sf, s...)), c.Info))
}
func Warn(s ...interface{}) {
	fmt.Println(l.time(), l.line()[0], l.color(l.ip(l.string(s...)), c.Warn))
}
func WarnF(sf string, s ...interface{}) {
	fmt.Println(l.time(), l.line()[0], l.color(l.ip(fmt.Sprintf(sf, s...)), c.Warn))
}
func Error(s ...interface{}) error {
	l.error(l.line(), s...)
	return errors.New(l.string(s...))
}
func Fatal(s ...interface{}) {
	l.error(l.line(), s...)
	os.Exit(0)
}

func Struct(s interface{}) {
	fmt.Println(l.time(), l.line()[0], l.color(s, c.Info))
}
func Json(s interface{}) {
	marshal, err := json.MarshalIndent(s, "", "\t") // 转格式化后json
	if err != nil {
		l.error(l.line(), s)
		return
	}
	Info("\n" + string(marshal))
}
func Progress(min int64, max int64, char string) {
	percentage := int(float64(min) / float64(max) * 100)
	var _char string
	for i := 0; i < percentage/2; i++ {
		_char += char
	}
	fmt.Printf("\r进度: [%-50s] %d%%   %d/%d", _char, percentage, min, max)
	
	if percentage == 100 {
		fmt.Println()
	}
}
