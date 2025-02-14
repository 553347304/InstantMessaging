package method

type stringServerInterface interface {
	Slice(int) string // 返回指定长度字符串 超出边界全部
}
type stringServer struct {
	String string
}

func String(s string) stringServerInterface {
	return &stringServer{String: s}
}
func (s *stringServer) Slice(length int) string {
	runes := []rune(s.String)
	if len(runes) < 4 {
		return s.String
	}
	return string(runes[:length])
}
