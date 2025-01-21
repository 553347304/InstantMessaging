package log_service

import (
	"bytes"
	"context"
	"fim_server/config/core"
	"fim_server/models/log_models"
	"fim_server/models/user_models"
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/method"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

type PusherServerInterface interface {
	Response(http.ResponseWriter, *http.Request, []byte) string
	Info(context.Context, string)
	Warn(context.Context, string)
	Error(context.Context, string)
}
type pusherServer struct {
	DB      *gorm.DB
	UserId  string `json:"user_id"`
	IP      string `json:"ip"`
	Type    string `json:"type"`
	Service string `json:"service"`
	method  string
	url     string
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewPusher(serviceName, _type string) PusherServerInterface {
	var c Config
	config, _ := ioutil.ReadFile("../../service.yaml")
	yaml.Unmarshal(config, &c)
	return &pusherServer{
		Service: serviceName,
		Type:    _type,
		DB:      core.Mysql(c.System.Mysql),
	}
}

func (p *pusherServer) Response(w http.ResponseWriter, r *http.Request, body []byte) string {
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
func (p *pusherServer) Info(ctx context.Context, content string) {
	go p.Save(ctx, "info", content)
}
func (p *pusherServer) Error(ctx context.Context, content string) {
	go p.Save(ctx, "error", content)
}
func (p *pusherServer) Warn(ctx context.Context, content string) {
	go p.Save(ctx, "warn", content)
}
func (p *pusherServer) Save(ctx context.Context, level, content string) {
	if p.Service == "log" && p.method == "GET" {
		return
	}
	p.UserId = ctx.Value("user_id").(string)
	if p.UserId == "" {
		return
	}
	p.IP = ctx.Value("ip").(string)
	addr := method.Addr().GetAddr(p.IP)
	userId := conv.Type(p.UserId).Uint()
	var user user_models.UserModel
	mutex := sync.Mutex{}
	mutex.Lock() // 加锁
	p.DB.Take(&user, userId)
	p.DB.Create(&log_models.LogModel{
		UserId:  userId,
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
