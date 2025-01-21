package logs

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func String(_color string, s ...interface{}) (string, string, string) {
	_source := ""
	for i := 0; i < len(s); i++ {
		_source += " " + fmt.Sprint(s[i])
	}
	_source = _source[1:]
	_string := logColor(isFieldColor(_source), _color)
	_time := logColor(fmt.Sprintf("[%v]", time.Now().Format(times)), "#ffffff")
	return _time, _string, _source
}

func Info(s ...interface{}) string {
	_time, _string, _source := String("#80ffff", s...)
	fmt.Println(_time, getLine()[line], _string)
	return _source
}
func InfoF(sf string, s ...interface{}) string {
	_time, _string, _source := String("#80ffff", fmt.Sprintf(sf, s...))
	fmt.Println(_time, getLine()[line], _string)
	return _source
}
func Error(s ...interface{}) error {
	_time, _string, _source := String("#F75464", s...)
	fmt.Println(_time, _string)
	row := getLine()
	for i := line; i < len(row); i++ {
		fmt.Println(row[i])
	}
	return errors.New(_source)
}
func Fatal(s ...interface{}) {
	_time, _string, _ := String("#F75464", s...)
	fmt.Println(_time, getLine()[line], _string)
	os.Exit(1)
}
