package service_method

import (
	"bytes"
	"context"
	"fim_server/models/log_models"
	"fim_server/models/user_models"
	"fim_server/utils/open_api/open_api_info"
	"fim_server/utils/src"
	"fim_server/utils/stores/conv"
	"fmt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"sync"
)

type ServerInterfaceLog interface {
	Response(http.ResponseWriter, *http.Request, []byte) string
	Info(context.Context, string)
	Warn(context.Context, string)
	Error(context.Context, string)
}
type serverLog struct {
	DB      *gorm.DB
	UserId  string `json:"user_id"`
	Ip      string `json:"ip"`
	Type    string `json:"type"`
	Service string `json:"service"`
	method  string
	url     string
}

//goland:noinspection GoExportedFuncWithUnexportedType
func Log(serviceName string, mode int) ServerInterfaceLog {
	maps := map[int]string{
		1: "登录日志",
		2: "操作日志",
		3: "运行日志",
	}
	return &serverLog{
		Service: serviceName,
		Type:    maps[mode],
		DB:      src.Client().Mysql("127.0.0.1:3306 baiyin fim_db"),
	}
}

func (p *serverLog) Response(w http.ResponseWriter, r *http.Request, body []byte) string {
	p.method = r.Method
	p.url = r.URL.String()
	header := string(conv.Json().Marshal(r.Header))
	request, _ := io.ReadAll(r.Body)
	response := string(body)
	r.Body = io.NopCloser(bytes.NewBuffer(request)) // 重新写入数据
	return fmt.Sprintf(""+
		`<div class="code">
			<div class="request">
				<span class="method">%s</span>
				<span class="url">%s</span>
				<pre class="header">%s</pre>
				<pre class="body">%s</pre>
			</div>
			<div class="response">
				<pre class="body">%s</pre>
			</div>
		</div>`, p.method, p.url, header, request, response)
}
func (p *serverLog) Info(ctx context.Context, content string) {
	go p.Save(ctx, "info", content)
}
func (p *serverLog) Error(ctx context.Context, content string) {
	go p.Save(ctx, "error", content)
}
func (p *serverLog) Warn(ctx context.Context, content string) {
	go p.Save(ctx, "warn", content)
}
func (p *serverLog) Save(ctx context.Context, level, content string) {
	if p.Service == "log" && p.method == "GET" {
		return
	}
	p.UserId = ctx.Value("user_id").(string)
	if p.UserId == "" {
		return
	}
	
	p.Ip = ctx.Value("ip").(string)
	addr := open_api_info.GetAddrByIP(p.Ip)
	UserId, err := conv.Type(p.UserId).Uint64()
	if err != nil {
		return
	}
	var user user_models.UserModel
	mutex := sync.Mutex{}
	mutex.Lock() // 加锁
	p.DB.Take(&user, UserId)
	p.DB.Create(&log_models.LogModel{
		UserId:   UserId,
		Username: user.Username,
		Avatar:   user.Avatar,
		Ip:       p.Ip,
		Addr:     addr,
		Service:  p.Service,
		Type:     p.Type,
		Level:    level,
		Content:  content,
	})
	mutex.Unlock() // 解锁
}
