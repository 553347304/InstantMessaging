package valid

type bcryptServiceInterface interface {
	Hash(string) string        // 加密密码	value
	Check(string, string) bool // 验证密码   hash | value
}
type imageCodeServiceInterface interface {
	ImageView() imageCodeResponse // 生成图片验证码
	Check(string, string) bool    // 验证   id | value
}
type jwtServiceInterface interface {
	Hash(interface{}) string // 加密 token
	Parse(string) *claims    // 解析 token
}
type md5ServiceInterface interface {
	Hash(string) string        // 加密	value
	Check(string, string) bool // 验证   hash | value
}

type bcryptService struct{}
type imageCodeService struct{}
type jwtService struct {
	key     string // 秘钥
	expires int    // 过期时间  /小时
	issuer  string // 签发人
}
type md5Service struct{}

func Bcrypt() bcryptServiceInterface       { return &bcryptService{} }
func ImageCode() imageCodeServiceInterface { return &imageCodeService{} }
func Jwt() jwtServiceInterface             { return &jwtService{key: "key", expires: 480, issuer: "baiyin"} }
func MD5() md5ServiceInterface             { return &md5Service{} }
