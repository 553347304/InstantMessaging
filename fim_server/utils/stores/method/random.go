package method

import "math/rand"

type serverInterfaceRandom interface {
	String(int) string // 返回指定长度字符串 超出边界全部
}

type serverRandom struct {
	English string
	Number  string
}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func Random() serverInterfaceRandom {
	return &serverRandom{
		English: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Number:  "0123456789",
	}
}

func (s *serverRandom) String(length int) string {
	char := []rune(s.Number + s.English)
	num := make([]rune, length)
	for i := range num {
		num[i] = char[rand.Intn(len(char))]
	}
	return string(num)
}
