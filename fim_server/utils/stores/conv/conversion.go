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
	Int() int   // 转换错误:-1
	Uint() uint // 转换错误:-1
	Error() error
}
type serverType struct{ Value string }

//goland:noinspection GoExportedFuncWithUnexportedType
func Type[T cmp.Ordered](value T) serverInterfaceType {
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
	return uint(s.Int())
}
func (s *serverType) Error() error {
	return errors.New(s.Value)
}
