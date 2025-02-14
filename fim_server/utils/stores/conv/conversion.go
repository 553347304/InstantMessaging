package conv

import (
	"cmp"
	"errors"
	"fim_server/utils/stores/logs"
	"fmt"
	"strconv"
)

type serverInterfaceType interface {
	String() string
	Uint() uint // 转换错误:0
	Uint64() (uint64, error)
	Int() int         // 转换错误:-1
	Float64() float64 // 转换错误:-1
	Error() error
}
type serverType struct{ Value string }

//goland:noinspection GoExportedFuncWithUnexportedType
func Type[T cmp.Ordered | any](value T) serverInterfaceType {
	return &serverType{Value: fmt.Sprint(value)}
}

func (s *serverType) String() string {
	return s.Value
}
func (s *serverType) Int() int {
	number, err := strconv.Atoi(s.Value)
	if err != nil {
		logs.Error("转换错误: ", err)
		return -1
	}
	return number
}
func (s *serverType) Uint() uint {
	number, err := strconv.Atoi(s.Value)
	if err != nil {
		logs.Error("转换错误: ", err)
		return 0
	}
	return uint(number)
}
func (s *serverType) Uint64() (uint64, error) {
	number, err := strconv.Atoi(s.Value)
	if err != nil {
		return 0, logs.Error("转换错误: ", err)
	}
	return uint64(number),nil
}
func (s *serverType) Float64() float64 {
	rounded, err := strconv.ParseFloat(s.Value, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
		return -1
	}
	return rounded
}
func (s *serverType) Error() error {
	if s.Value == "" {
		return nil
	}
	return errors.New(s.Value)
}
