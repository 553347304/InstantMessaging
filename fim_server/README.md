# fim_server

## 即时通讯后端

### 命令行参数

``` shell
go run main.go -db # 迁移表结构
```

### service

``` yaml
cd service
$t=@("-style", "go_zero", "--home", "template");

$p="user";      # 用户
$p="chat";      # 消息
$p="auth";      # 校验
$p="file";      # 文件
$p="setting";   # 设置
$p="group";     # 群聊
$p="log";       # 日志
$p="test";       # 日志

$v="rpc/$p"; goctl rpc protoc "$v.proto" --go_out=$v --zrpc_out=$v -m --go-grpc_out=$v $t -m   # 生成RPC
$v="api/$p"; goctl api go -api "$v.api" -dir "$v" $t                                           # 生成API
```

#### 解决卡进程

``` cmd
rmdir /s /q C:\Users\baiyin\AppData\Local\JetBrains\IntelliJIdea2024.3\tmp\GoLand
```

