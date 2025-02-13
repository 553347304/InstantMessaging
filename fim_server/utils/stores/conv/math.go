package conv

type mathServerInterface interface {
	Byte(int64, int64) float64
}
type mathServer struct{}

func Math() mathServerInterface { return &mathServer{} }

func (*mathServer) Byte(b int64, new int64) float64 {
	return float64(b) / float64(new)
}
