# fim_server

## 即时通讯后端

### 命令行参数
``` shell
go run main.go -db # 迁移表结构
```

### service
``` yaml
goctl template init --home $pwd/go-zero-template    # 模版文件

cd service
$t=@("-style", "go_zero", "--home", "template");

$p="user";      # 用户
$p="chat";      # 消息
$p="auth";      # 校验
$p="file";      # 文件
$p="setting";   # 设置
$p="group";     # 群聊

$v="rpc/$p"; goctl rpc protoc "$v.proto" --go_out=$v --zrpc_out=$v --go-grpc_out=$v $t   # 生成RPC
$v="api/$p"; goctl api go -api "$v.api" -dir "$v" $t                                     # 生成API
```

## 模块
``` yaml
go get gopkg.in/yaml.v3                                             # yaml
go get gorm.io/gorm                                                 # gorm
go get gorm.io/driver/mysql                                         # mysql
go get github.com/go-redis/redis                                    # redis
go get github.com/gorilla/websocket                                 # websocket
go get go.etcd.io/etcd/client/v3                                    # etcd
go get -u github.com/zeromicro/go-zero                              # go-zero
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0     # protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0     # protoc-gen-go-grpc
# go-zero protoc
curl -o "$env:USERPROFILE\go\protoc.zip" https://github.com/protocolbuffers/protobuf/releases/download/v3.9.0/protoc-3.9.0-win64.zip
7z x "$env:USERPROFILE\go\protoc.zip" -o"$env:USERPROFILE\go\" -y
rm "$env:USERPROFILE\go\protoc.zip"
protoc --version
# go-zero goctl
goctl template init --home template     # 指定模版文件
goctl api go -api main.api -dir .       # 生成api文件
goctl rpc protoc main.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.     # 生成rpc文件
```

## utils
``` yaml
go get -u golang.org/x/crypto/bcrypt    # bcrypt
go get github.com/golang-jwt/jwt/v4     # jwt
```

### 初始化
``` yaml
go mod init main    # 初始化模块
go clean -modcache  # 清理模块缓存
go mod tidy         # 重新加载依赖
go mod graph        # 查看依赖关系
go env -w GOPROXY=https://goproxy.cn,direct # 环境
go env GOPROXY
```