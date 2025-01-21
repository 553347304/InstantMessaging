package log_service

const (
	Login  = "登录日志"
	Action = "操作日志"
)

type Config struct {
	System struct {
		Mysql string `yaml:"Mysql"`
		Redis string `yaml:"Redis"`
		Etcd  string `yaml:"Etcd"`
	} `yaml:"System"`
}
