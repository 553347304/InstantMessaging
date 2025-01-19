package method

import "time"

type serverInterfaceTime interface {
	Now() string // 返回指定长度字符串 超出边界全部
}

type serverTime struct {
	Format string
}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Time() serverInterfaceTime {
	return &serverTime{
		Format: "2006-01-02 15:04:05",
	}
}

func (s *serverTime) Now() string {
	return time.Now().Format(s.Format)
}
