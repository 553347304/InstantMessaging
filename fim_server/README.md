# fim_server

# 即时通讯后端

### 命令行参数
```Command
go run main.go -db # 迁移表结构
```

### 模块
``` yaml
go get gopkg.in/yaml.v3
```

```gorm mysql
go get gorm.io/gorm
go get gorm.io/driver/mysql
```

```zero
go install github.com/zeromicro/go-zero/tools/goctl@latest
```
### zero模版
```goctl
goctl template init --home template
```

### 环境
```
go env -w GOPROXY=https://goproxy.cn,direct
go env GOPROXY
```

### server
```
goctl api go -api auth_api.api -dir .
=
```

