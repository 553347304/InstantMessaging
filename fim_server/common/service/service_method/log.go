package service_method

import (
	"bytes"
	"context"
	"fim_server/models/log_models"
	"fim_server/models/user_models"
	"fim_server/utils/src"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/method"
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
	UserID  string `json:"user_id"`
	IP      string `json:"ip"`
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
		DB:      src.Client().Mysql("root:baiyin@tcp(127.0.0.1:3306)/fim_db?charset=utf8mb4&parseTime=True&loc=Local"),
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
	p.UserID = ctx.Value("user_id").(string)
	if p.UserID == "" {
		return
	}
	p.IP = ctx.Value("ip").(string)
	addr := method.Addr().GetAddr(p.IP)
	UserID := conv.Type(p.UserID).Uint()
	var user user_models.UserModel
	mutex := sync.Mutex{}
	mutex.Lock() // 加锁
	p.DB.Take(&user, UserID)
	p.DB.Create(&log_models.LogModel{
		UserID:  UserID,
		Name:    user.Name,
		Avatar:  user.Avatar,
		IP:      p.IP,
		Addr:    addr,
		Service: p.Service,
		Type:    p.Type,
		Level:   level,
		Content: content,
	})
	mutex.Unlock() // 解锁
}
