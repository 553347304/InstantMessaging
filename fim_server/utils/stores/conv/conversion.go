package conv

import (
	"cmp"
	"errors"
	"fim_server/utils/stores/logs"
	"fmt"
	"strconv"
)

func HexRgb(hex string) (int, int, int) {
	if hex[0] == '#' {
		hex = hex[1:]
	}
	if len(hex) != 6 {
		return 0, 0, 0
	}
	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)
	return int(r), int(g), int(b)
}

type typeServerInterface interface {
	String() string
	Int() int   // 转换错误返回   -1
	Uint() uint // 转换错误返回   -1
	Error() error
}
type typeServer struct {
	Value string
}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Type[T cmp.Ordered](value T) typeServerInterface {
	return &typeServer{Value: fmt.Sprint(value)}
}

func (s *typeServer) String() string {
	return s.Value
}
func (s *typeServer) Int() int {
	number, err := strconv.Atoi(s.Value)
	if err != nil {
		logs.Error("转换错误: ", err)
		return -1
	}
	return number
}
func (s *typeServer) Uint() uint {
	return uint(s.Int())
}
func (s *typeServer) Error() error {
	return errors.New(s.Value)
}
