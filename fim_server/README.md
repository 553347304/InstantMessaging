# fim_server

## 即时通讯后端

### 命令行参数
```command
go run main.go -db # 迁移表结构
```

### server
```zero
goctl template init --home template     # 指定模版文件
goctl api go -api main.api -dir .
```





### 模块
``` yaml
go get gopkg.in/yaml.v3
```
```gorm mysql
go get gorm.io/gorm
go get gorm.io/driver/mysql
```
```redis
github.com/go-redis/redis
```
```zero rpc
go get -u github.com/zeromicro/go-zero
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
curl -o "$env:USERPROFILE\go\protoc.zip" https://github.com/protocolbuffers/protobuf/releases/download/v3.9.0/protoc-3.9.0-win64.zip
7z x "$env:USERPROFILE\go\protoc.zip" -o"$env:USERPROFILE\go\" -y
rm "$env:USERPROFILE\go\protoc.zip"
protoc --version
```




### utils
```bcrypt
go get -u golang.org/x/crypto/bcrypt
```
```jwt
go get github.com/golang-jwt/jwt/v4
```



### 环境
```
go env -w GOPROXY=https://goproxy.cn,direct
go env GOPROXY
```